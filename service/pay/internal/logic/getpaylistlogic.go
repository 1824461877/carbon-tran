package logic

import (
	"context"
	"errors"

	"pay/internal/svc"
	"pay/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayListLogic {
	return &GetPayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPayListLogic) GetPayList(in *pb.PayOnceReq) (*pb.PayListResp, error) {
	// todo: add your logic here and delete this line
	uid, ok := l.ctx.Value("uid").(string)
	if uid == "" && !ok {
		return nil, errors.New("uid is empty")
	}

	payOrder, err := l.svcCtx.PayOrder.FindOneByPayOrderId(l.ctx, in.PayOrderId)
	if err != nil {
		return nil, err
	}

	if uid != payOrder.Initiator {
		return &pb.PayListResp{
			PayListResp: []*pb.PayOnceResp{},
		}, nil
	}

	return &pb.PayListResp{
		PayListResp: []*pb.PayOnceResp{
			{
				PayOrderId:    payOrder.PayOrderId,
				Initiator:     payOrder.Initiator,
				Recipient:     payOrder.Recipient,
				PayStatus:     int32(payOrder.PayStatus),
				PayAmount:     payOrder.PayAmount,
				InitiatorTime: payOrder.InitiatorTime.String(),
				FinishTime:    payOrder.FinishTime.String(),
			},
		},
	}, nil
}
