package logic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	pb2 "pay/pb"
	"time"
	"trade/internal/middleware"
	"trade/internal/svc"
	"trade/internal/types"
	"trade/model"
	"trade/pb"
	"trade/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type TradeExecutionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTradeExecutionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TradeExecutionLogic {
	return &TradeExecutionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TradeExecutionLogic) PayExecution(po types.PayOrder) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, po)
	return token.SignedString([]byte(l.svcCtx.Config.ServiceJwtSign.PayServiceAuth.JwtSignKey))
}

func (l *TradeExecutionLogic) TradeExecution(in *pb.TradeReq) (*pb.TradeExecutionResp, error) {
	var (
		claims    *types.TradeOrder
		token     *jwt.Token
		ve        *jwt.ValidationError
		PayToken  string
		execution *pb2.PayExecutionResp
		err       error
		as        *model.AssetsSell
		ok        bool
	)

	token, err = jwt.ParseWithClaims(in.TaskToken, &types.TradeOrder{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.ServiceJwtSign.TradeServiceAuth.JwtSignKey), nil
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

	if claims, ok = token.Claims.(*types.TradeOrder); !ok && !token.Valid {
		return nil, errors.New("token is error")
	}

	as, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByExId(l.ctx, claims.ExchangeAssetID)
	if err != nil {
		return nil, err
	}

	if PayToken, err = l.PayExecution(types.PayOrder{
		PayAmount:    claims.PayAmount,
		Initiator:    claims.Initiator,
		CollectionID: as.CollectionWalletId,
		Recipient:    claims.Recipient,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(l.svcCtx.Config.ServiceJwtSign.PayServiceAuth.JwtSignExpire) * time.Minute).Unix(),
		},
	}); err != nil {
		return nil, err
	}

	times := time.Now()
	execution, err = l.svcCtx.ServiceRpc.PayRpc.PayExecution(l.ctx, &pb2.PayReq{
		TaskToken: PayToken,
		ReqTime:   times.String(),
	})
	if err != nil {
		return nil, err
	}

	TID := utils.TID()
	_, err = l.svcCtx.MysqlServiceContext.TradeOrder.Insert(l.ctx, &model.TradeOrder{
		TradeOrderId:    TID,
		ExchangeAssetID: claims.ExchangeAssetID,
		PayOrderId:      execution.PayOrderId,
		CarbonAssetId:   as.AssId,
		CollectionID:    claims.CollectionID,
		Initiator:       claims.Initiator,
		Recipient:       claims.Recipient,
		Number:          claims.Number,
		TradeStatus:     types.Incomplete,
		InitiatorTime:   times,
		FinishTime:      times,
	})
	if err != nil {
		return nil, err
	}

	return &pb.TradeExecutionResp{
		TradeOrderId: execution.PayOrderId,
		Status:       int32(types.Incomplete),
		CreateTime:   times.String(),
	}, nil
}
