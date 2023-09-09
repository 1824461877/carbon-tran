package logic

import (
	"context"
	"trade/internal/svc"
	"trade/model"
	"trade/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradeObtainListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTradeObtainListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradeObtainListLogic {
	return &GetTradeObtainListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTradeObtainListLogic) GetTradeObtainList(in *pb.TradeObtainReq) (*pb.TradeListResp, error) {
	var (
		tradeOrder *[]model.TradeOrder
		err        error
	)

	switch in.State {
	case "initiator":
		tradeOrder, err = l.svcCtx.MysqlServiceContext.TradeOrder.FindOneByInitiatorTradeOrderIdList(l.ctx, in.Obtain)
		if err != nil {
			return nil, err
		}
	case "recipient":
		tradeOrder, err = l.svcCtx.MysqlServiceContext.TradeOrder.FindOneByRecipientTradeOrderIdList(l.ctx, in.Obtain)
		if err != nil {
			return nil, err
		}
	}

	var r []*pb.TradeOnceResp
	if tradeOrder != nil {
		for _, v := range *tradeOrder {
			r = append(r, &pb.TradeOnceResp{
				TradeOrderId:    v.TradeOrderId,
				CarbonAssetId:   v.CarbonAssetId,
				ExchangeAssetId: v.ExchangeAssetID,
				PayOrderId:      v.PayOrderId,
				Initiator:       v.Initiator,
				Number:          v.Number,
				Recipient:       v.Recipient,
				TadeStatus:      v.TradeStatus,
				InitiatorTime:   v.InitiatorTime.Format("2006-01-02-15:04:05"),
				FinishTime:      v.FinishTime.Format("2006-01-02-15:04:05"),
			})
		}
	}

	return &pb.TradeListResp{
		TradeOnceResp: r,
	}, nil
}
