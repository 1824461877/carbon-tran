package logic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
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

func (l *TradeExecutionLogic) TradeExecution(in *pb.TradeReq) (*pb.TradeExecutionResp, error) {
	// todo: add your logic here and delete this line
	var (
		claims *types.TradeOrder
		token  *jwt.Token
		ve     *jwt.ValidationError
		err    error
		ok     bool
	)
	token, err = jwt.ParseWithClaims(in.TaskToken, &types.TradeOrder{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.TradeServiceAuth.JwtSignKey), nil
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
	times := time.Now()
	TID := utils.TID()
	_, err = l.svcCtx.TradeOrder.Insert(l.ctx, &model.TradeOrder{
		TradeOrderId:  TID,
		PayOrderId:    claims.PayOrderId,
		CarbonAssetId: claims.CarbonAssetId,
		Initiator:     claims.Initiator,
		Recipient:     claims.Recipient,
		TadeStatus:    int64(types.Incomplete),
		InitiatorTime: times,
		FinishTime:    times,
	})
	if err != nil {
		return nil, err
	}

	return &pb.TradeExecutionResp{
		TradeOrderId: TID,
		Status:       types.Incomplete,
		CreateTime:   times.String(),
	}, nil
}
