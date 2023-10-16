package logic

import (
	"context"
	"trade/internal/types"
	"trade/model"

	"trade/internal/svc"
	"trade/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllTradeOrderObtainListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllTradeOrderObtainListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTradeOrderObtainListLogic {
	return &GetAllTradeOrderObtainListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllTradeOrderObtainListLogic) GetAllTradeOrderObtainList(in *pb.GetAllTradeOrderObtainReq) (*pb.TradeListResp, error) {
	var (
		limit      int32
		offset     int32
		tradeOrder *[]model.TradeOrder
		err        error
	)

	if in.Search != "" {
		tradeOrder, err = l.svcCtx.MysqlServiceContext.TradeOrder.FindTradeSearch(l.ctx, in.Search)
		//if err != nil && err != sqlc.ErrNotFound {
		//	return nil, err
		//}
	} else {
		limit = in.Limit
		if in.Pages <= 0 {
			in.Pages = 1
		}
		offset = (in.Pages - 1) * in.Limit
		tradeOrder, err = l.svcCtx.MysqlServiceContext.TradeOrder.FindTradeStatus(l.ctx, limit, offset, types.Completed)
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

	var (
		count *int64
	)
	count, err = l.svcCtx.MysqlServiceContext.TradeOrder.FindPageCount(l.ctx)
	if err != nil {
		return nil, err
	}

	return &pb.TradeListResp{
		TradeOnceResp: r,
		PageCount:     *count,
	}, nil
}
