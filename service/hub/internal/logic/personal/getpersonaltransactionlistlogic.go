package personal

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	pb2 "trade/pb"
)

type GetPersonalTransactionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalTransactionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalTransactionListLogic {
	return &GetPersonalTransactionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalTransactionListLogic) GetPersonalTransactionList(uid string, req *types.GetObtainTransactionReq) (resp *types.ObtainTransactionResp, err error) {
	var (
		tradeList *pb2.TradeListResp
	)

	tradeList, err = l.svcCtx.ServiceRpc.TradeRpc.GetTradeObtainList(l.ctx, &pb2.TradeObtainReq{
		Obtain: uid,
		State:  req.State,
	})
	if err != nil {
		return nil, err
	}

	return &types.ObtainTransactionResp{
		Code:          types.SuccessCode,
		TradeOnceResp: tradeList.TradeOnceResp,
	}, nil
}
