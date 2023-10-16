package record

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"pay/pb"
)

type PayOnceRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPayOnceRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOnceRecordLogic {
	return &PayOnceRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOnceRecordLogic) PayOnceRecord(req *types.PayOnceRecordReq) (resp *types.PayOnceRecordResp, err error) {
	var (
		payOnceList *pb.PayOnceList
	)

	payOnceList, err = l.svcCtx.ServiceRpc.PayRpc.GetBasicPayOnce(l.ctx, &pb.PayOnceReq{
		PayOrderId: req.PayOrderId,
	})
	if err != nil {
		return nil, err
	}

	return &types.PayOnceRecordResp{
		PayOnceList: payOnceList,
		Code:        types.SuccessCode,
	}, nil
}
