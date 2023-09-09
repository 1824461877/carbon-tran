package logic

import (
	"context"
	"encoding/json"
	"github.com/plutov/paypal"
	"github.com/zeromicro/go-zero/core/logx"
	"pay/internal/svc"
	"pay/internal/types"
	"pay/model"
	"pay/pb"
	pb2 "trade/pb"
)

type PayApproveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayApproveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayApproveLogic {
	return &PayApproveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayApproveLogic) PayApprove(in *pb.PayApproveReq) (*pb.PayApproveResp, error) {
	// todo: add your logic here and delete this line
	var (
		err error
		//approve     *paypal.CaptureOrderResponse
		get string
		//orderUpdate *pb2.TradeOrderUpdateResp
	)

	get, err = l.svcCtx.Redis.Get(in.PayOrderId)
	if err != nil {
		return nil, err
	}
	var po = &model.PayOrder{}
	err = json.Unmarshal([]byte(get), po)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.PayPalService.GetToken()
	if err != nil {
		return nil, err
	}

	_, err = l.PayOrderApprove(l.svcCtx.PayPalService.Client, po.PayId)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.ServiceRpc.TradeRpc.TradeOrderUpdate(l.ctx, &pb2.TradeOrderUpdateReq{
		PayOrderId:  in.PayOrderId,
		UpdateState: int64(types.Paid),
	})
	if err != nil {
		return nil, err
	}

	return &pb.PayApproveResp{
		PayStatus: "ok",
	}, nil
}

func (l *PayApproveLogic) PayOrderApprove(pay *paypal.Client, orderID string) (*paypal.CaptureOrderResponse, error) {
	return pay.CaptureOrder(orderID, paypal.CaptureOrderRequest{
		PaymentSource: &paypal.PaymentSource{},
	})
}

func (l *PayApproveLogic) PayGetOrder(pay *paypal.Client, orderID string) (*paypal.Order, error) {
	return pay.GetOrder(orderID)
}
