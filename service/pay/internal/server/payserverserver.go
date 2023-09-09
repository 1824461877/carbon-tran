// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package server

import (
	"context"

	"pay/internal/logic"
	"pay/internal/svc"
	"pay/pb"
)

type PayServerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPayServerServer
}

func NewPayServerServer(svcCtx *svc.ServiceContext) *PayServerServer {
	return &PayServerServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServerServer) GetPayList(ctx context.Context, in *pb.PayOnceReq) (*pb.PayListResp, error) {
	l := logic.NewGetPayListLogic(ctx, s.svcCtx)
	return l.GetPayList(in)
}

func (s *PayServerServer) PayExecution(ctx context.Context, in *pb.PayReq) (*pb.PayExecutionResp, error) {
	l := logic.NewPayExecutionLogic(ctx, s.svcCtx)
	return l.PayExecution(in)
}

func (s *PayServerServer) PayOrderStatus(ctx context.Context, in *pb.PayOrderStatusReq) (*pb.PayOrderStatusResp, error) {
	l := logic.NewPayOrderStatusLogic(ctx, s.svcCtx)
	return l.PayOrderStatus(in)
}

func (s *PayServerServer) PayApprove(ctx context.Context, in *pb.PayApproveReq) (*pb.PayApproveResp, error) {
	l := logic.NewPayApproveLogic(ctx, s.svcCtx)
	return l.PayApprove(in)
}
