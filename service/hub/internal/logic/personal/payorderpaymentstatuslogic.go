package personal

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"pay/pb"
)

type PayOrderPaymentStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderPaymentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderPaymentStatusLogic {
	return &PayOrderPaymentStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderPaymentStatusLogic) PayOrderPaymentStatus(req *types.PayOrderReq) (resp *types.PayOrderStatusResp, err error) {
	status, err := l.svcCtx.ServiceRpc.PayRpc.PayOrderStatus(l.ctx, &pb.PayOrderStatusReq{
		PayOrderId: req.PayOrderId,
	})
	if err != nil {
		return nil, err
	}

	return &types.PayOrderStatusResp{
		Code:               types.SuccessCode,
		PayOrderStatusResp: status,
	}, nil
}
