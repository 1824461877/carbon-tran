package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/plutov/paypal"
	"pay/internal/middleware"
	"pay/internal/types"
	"pay/model"
	"pay/utils"
	"time"

	"pay/internal/svc"
	"pay/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayExecutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type PayDo struct {
	Value    string
	Currency string
	CID      string
}

func NewPayExecutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayExecutionLogic {
	return &PayExecutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//func (l *PayExecutionLogic) PayExecution(in *proto.PayReq) (*proto.PayExecutionResp, error) {
//	todo: add your logic here and delete this line
//	var (
//		claims *types.PayOrder
//		token  *jwt.Token
//		ve     *jwt.ValidationError
//		err    error
//		ok     bool
//	)
//
//	token, err = jwt.ParseWithClaims(in.TaskToken, &types.PayOrder{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(l.svcCtx.Config.PayServiceAuth.JwtSignKey), nil
//	})
//
//	if err != nil {
//		if ve, ok = err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				return nil, middleware.TokenMalformed
//			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
//				return nil, middleware.TokenExpired
//			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
//				return nil, middleware.TokenNotValidYet
//			} else {
//				return nil, middleware.TokenInvalid
//			}
//		}
//	}
//
//	if claims, ok = token.Claims.(*types.PayOrder); !ok && !token.Valid {
//		return nil, errors.New("token is error")
//	}
//
//	times := time.Now()
//	var PayStatus = types.Unpaid
//	if err = l.PayWallet(claims.InitiatorWalletId, claims.RecipientWalletId, claims.PayAmount); err == nil {
//		PayStatus = types.Paid
//	}
//
//	var PID = utils.PID()
//	if _, err = l.svcCtx.PayOrder.Insert(l.ctx, &model.PayOrder{
//		PayOrderId:    PID,
//		Initiator:     claims.Initiator,
//		Recipient:     claims.Recipient,
//		PayStatus:     int64(PayStatus),
//		PayAmount:     int64(claims.PayAmount),
//		InitiatorTime: times,
//	}); err != nil {
//		return nil, err
//	}
//
//	return &proto.PayExecutionResp{
//		PayOrderId: PID,
//		PayStatus:  PayStatus,
//		CreateTime: times.String(),
//	}, nil
//}

func (l *PayExecutionLogic) PayExecution(in *pb.PayReq) (*pb.PayExecutionResp, error) {
	// todo: add your logic here and delete this line
	var (
		claims     *types.PayOrder
		token      *jwt.Token
		ve         *jwt.ValidationError
		UserWallet *model.UserWallet
		err        error
		ok         bool
	)

	token, err = jwt.ParseWithClaims(in.TaskToken, &types.PayOrder{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.PayServiceAuth.JwtSignKey), nil
	})

	if err != nil {
		if ve, ok = err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, middleware.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, middleware.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, middleware.TokenNotValidYet
			} else {
				return nil, middleware.TokenInvalid
			}
		}
	}

	if claims, ok = token.Claims.(*types.PayOrder); !ok && !token.Valid {
		return nil, errors.New("token is error")
	}

	UserWallet, err = l.svcCtx.UserWallet.FindOneByWalletId(l.ctx, claims.CollectionID)
	if err != nil {
		return nil, err
	}

	times := time.Now()
	var (
		PayStatus = types.Unpaid
		order     *paypal.Order
	)
	//if err = l.PayWallet(claims.InitiatorWalletId, claims.RecipientWalletId, claims.PayAmount); err == nil {
	//	PayStatus = types.Paid
	//}

	err = l.svcCtx.PayPalService.GetToken()
	if err != nil {
		return nil, err
	}

	if order, err = l.PayOrder(l.svcCtx.PayPalService.Client, &PayDo{
		Value:    fmt.Sprintf("%.2f", claims.PayAmount),
		Currency: "USD",
		CID:      UserWallet.CID,
	}); err != nil {
		return nil, err
	}

	var PID = utils.PID()
	var data []byte
	var po = &model.PayOrder{
		PayOrderId:    PID,
		Initiator:     claims.Initiator,
		PayId:         order.ID,
		Recipient:     claims.Recipient,
		PayStatus:     int64(PayStatus),
		PayAmount:     int64(claims.PayAmount),
		InitiatorTime: times,
		FinishTime:    times,
	}

	if data, err = json.Marshal(po); err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.Redis.SetnxExCtx(l.ctx, PID, string(data), 300); err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.PayOrder.Insert(l.ctx, &model.PayOrder{
		PayOrderId:    PID,
		Initiator:     claims.Initiator,
		PayId:         order.ID,
		Recipient:     claims.Recipient,
		PayStatus:     int64(PayStatus),
		PayAmount:     int64(claims.PayAmount),
		InitiatorTime: times,
		FinishTime:    times,
	}); err != nil {
		return nil, err
	}

	return &pb.PayExecutionResp{
		PayOrderId: PID,
		PayStatus:  PayStatus,
		CreateTime: times.String(),
	}, nil
}

func (l *PayExecutionLogic) PayOrder(pay *paypal.Client, do *PayDo) (*paypal.Order, error) {
	var (
		orderResponse *paypal.Order
		err           error
	)
	if orderResponse, err = pay.CreateOrder(
		paypal.OrderIntentCapture,
		[]paypal.PurchaseUnitRequest{
			{
				Payee: &paypal.PayeeForOrders{
					EmailAddress: do.CID,
				},
				Amount: &paypal.PurchaseUnitAmount{
					Value:    do.Value,
					Currency: do.Currency,
				},
			},
		},
		&paypal.CreateOrderPayer{},
		&paypal.ApplicationContext{},
	); err != nil {
		return nil, err
	}
	return orderResponse, nil
}

//func (l *PayExecutionLogic) PayWallet(initiatorWalletId string, recipientWalletId string, amount float64) error {
//	var (
//		initiator, recipient *model.UserWallet
//		err                  error
//	)
//	initiator, err = l.svcCtx.UserWallet.FindOneByWalletId(l.ctx, initiatorWalletId)
//	if err != nil {
//		return err
//	}
//
//	recipient, err = l.svcCtx.UserWallet.FindOneByWalletId(l.ctx, recipientWalletId)
//	if err != nil {
//		return err
//	}
//
//	if initiator.Amount >= amount {
//		initiator.Amount = initiator.Amount - amount
//		recipient.Amount = recipient.Amount + amount
//	} else {
//		return errors.New("insufficient wallet balance, payment failed")
//	}
//
//	if err = l.svcCtx.UserWallet.Update(l.ctx, initiator); err != nil {
//		return err
//	}
//	if err = l.svcCtx.UserWallet.Update(l.ctx, recipient); err != nil {
//		return err
//	}
//	return nil
//}
