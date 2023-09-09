package logic

import (
	"context"
	"encoding/json"
	"pay/internal/svc"
	"pay/internal/types"
	"pay/model"
	"pay/pb"
	pb2 "trade/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderStatusLogic {
	return &PayOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayOrderStatusLogic) PayOrderStatus(in *pb.PayOrderStatusReq) (*pb.PayOrderStatusResp, error) {
	// todo: add your logic here and delete this line
	var (
		ttl int
		get string
		err error
	)

	ttl, err = l.svcCtx.Redis.Ttl(in.PayOrderId)
	if err != nil {
		return nil, err
	}

	if ttl == 0 {
		if _, err = l.svcCtx.ServiceRpc.TradeRpc.TradeOrderUpdate(l.ctx, &pb2.TradeOrderUpdateReq{
			PayOrderId:  in.PayOrderId,
			UpdateState: int64(types.Expire),
		}); err != nil {
			return nil, err
		}
		return &pb.PayOrderStatusResp{
			Code:   400,
			PayID:  "",
			TTL:    int32(ttl),
			Status: "The order has expired or does not exist",
		}, nil
	}

	get, err = l.svcCtx.Redis.Get(in.PayOrderId)
	if err != nil {
		return nil, err
	}
	var po = &model.PayOrder{}
	err = json.Unmarshal([]byte(get), po)
	if err != nil {
		return nil, err
	}

	return &pb.PayOrderStatusResp{
		Code:   200,
		PayID:  po.PayId,
		TTL:    int32(ttl),
		Status: "Order not expired",
	}, nil
}
