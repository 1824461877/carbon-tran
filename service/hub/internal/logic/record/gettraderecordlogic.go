package record

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"trade/pb"
)

type TradeRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTradeRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TradeRecordLogic {
	return &TradeRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TradeRecordLogic) TradeRecord(req *types.TradeRecordReq) (resp *types.TradeRecordResp, err error) {
	var (
		list *pb.TradeListResp
	)

	//if req.Search != "search" && len(req.Search) < 7 {
	//	return nil, errors.New("error")
	//}

	list, err = l.svcCtx.ServiceRpc.TradeRpc.GetAllTradeOrderObtainList(l.ctx, &pb.GetAllTradeOrderObtainReq{
		Limit:  10,
		Search: req.Search,
		Pages:  req.Page,
	})
	if err != nil {
		return
	}

	return &types.TradeRecordResp{
		Code:          types.SuccessCode,
		TradeListData: list,
	}, nil
}
