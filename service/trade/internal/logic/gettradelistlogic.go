package logic

import (
	"context"
	"errors"
	"trade/internal/svc"
	"trade/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTradeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradeListLogic {
	return &GetTradeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTradeListLogic) GetTradeList(in *pb.TradeOrderIdReq) (*pb.TradeListResp, error) {
	// todo: add your logic here and delete this line
	uid, ok := l.ctx.Value("uid").(string)
	if uid == "" && !ok {
		return nil, errors.New("uid is empty")
	}

	tradeOrder, err := l.svcCtx.TradeOrder.FindOneByTradeOrderId(l.ctx, in.TradeOrderId)
	if err != nil {
		return nil, err
	}

	if uid != tradeOrder.Initiator {
		return &pb.TradeListResp{
			TradeOnceResp: []*pb.TradeOnceResp{},
		}, nil
	}

	return &pb.TradeListResp{
		TradeOnceResp: []*pb.TradeOnceResp{
			{
				TradeOrderId:  tradeOrder.TradeOrderId,
				PayOrderId:    tradeOrder.PayOrderId,
				CarbonAssetId: tradeOrder.CarbonAssetId,
				Initiator:     tradeOrder.Initiator,
				Recipient:     tradeOrder.Recipient,
				TadeStatus:    tradeOrder.TadeStatus,
				InitiatorTime: tradeOrder.InitiatorTime.String(),
				FinishTime:    tradeOrder.FinishTime.String(),
			},
		},
	}, nil
}
