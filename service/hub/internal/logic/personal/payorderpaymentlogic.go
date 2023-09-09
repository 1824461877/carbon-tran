package personal

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"pay/pb"
)

type PayOrderPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderPaymentLogic {
	return &PayOrderPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderPaymentLogic) PayOrderPayment(req *types.PayOrderReq) (resp *types.PayOrderResp, err error) {
	approve, err := l.svcCtx.ServiceRpc.PayRpc.PayApprove(l.ctx, &pb.PayApproveReq{
		PayOrderId: req.PayOrderId,
	})
	if err != nil {
		return nil, err
	}

	return &types.PayOrderResp{
		Code:    types.SuccessCode,
		Messing: fmt.Sprintf("%v, pay_order_id: %v", approve.PayStatus, req.PayOrderId),
	}, nil
}
