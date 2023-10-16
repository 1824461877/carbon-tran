package logic

import (
	"context"
	"pay/model"

	"pay/internal/svc"
	"pay/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBasicPayOnceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBasicPayOnceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBasicPayOnceLogic {
	return &GetBasicPayOnceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBasicPayOnceLogic) GetBasicPayOnce(in *pb.PayOnceReq) (*pb.PayOnceList, error) {
	var (
		payOrderOnce *model.PayOrder
		err          error
		//order        *paypal.Order
	)
	payOrderOnce, err = l.svcCtx.PayOrder.FindOneByPayOrderId(l.ctx, in.PayOrderId)
	if err != nil {
		return nil, err
	}

	//order, err = l.svcCtx.PayPalService.Client.GetOrder(payOrderOnce.PayId)
	//if err != nil {
	//	return nil, err
	//}

	//order.PurchaseUnits

	return &pb.PayOnceList{
		PayOrderId: payOrderOnce.PayOrderId,
		//PayId:         payOrderOnce.PayId,
		Initiator:     payOrderOnce.Initiator,
		Recipient:     payOrderOnce.Recipient,
		PayStatus:     payOrderOnce.PayStatus,
		PayAmount:     payOrderOnce.PayAmount,
		InitiatorTime: payOrderOnce.InitiatorTime.Format("2006/01/02 15:04:05"),
		FinishTime:    payOrderOnce.FinishTime.Format("2006/01/02 15:04:05"),
	}, nil
}
