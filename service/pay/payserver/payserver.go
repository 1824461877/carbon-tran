// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package payserver

import (
	"context"

	"pay/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PayApproveReq      = pb.PayApproveReq
	PayApproveResp     = pb.PayApproveResp
	PayExecutionResp   = pb.PayExecutionResp
	PayListResp        = pb.PayListResp
	PayOnceReq         = pb.PayOnceReq
	PayOnceResp        = pb.PayOnceResp
	PayOrderStatusReq  = pb.PayOrderStatusReq
	PayOrderStatusResp = pb.PayOrderStatusResp
	PayReq             = pb.PayReq

	PayServer interface {
		GetPayList(ctx context.Context, in *PayOnceReq, opts ...grpc.CallOption) (*PayListResp, error)
		PayExecution(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayExecutionResp, error)
		PayOrderStatus(ctx context.Context, in *PayOrderStatusReq, opts ...grpc.CallOption) (*PayOrderStatusResp, error)
		PayApprove(ctx context.Context, in *PayApproveReq, opts ...grpc.CallOption) (*PayApproveResp, error)
	}

	defaultPayServer struct {
		cli zrpc.Client
	}
)

func NewPayServer(cli zrpc.Client) PayServer {
	return &defaultPayServer{
		cli: cli,
	}
}

func (m *defaultPayServer) GetPayList(ctx context.Context, in *PayOnceReq, opts ...grpc.CallOption) (*PayListResp, error) {
	client := pb.NewPayServerClient(m.cli.Conn())
	return client.GetPayList(ctx, in, opts...)
}

func (m *defaultPayServer) PayExecution(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayExecutionResp, error) {
	client := pb.NewPayServerClient(m.cli.Conn())
	return client.PayExecution(ctx, in, opts...)
}

func (m *defaultPayServer) PayOrderStatus(ctx context.Context, in *PayOrderStatusReq, opts ...grpc.CallOption) (*PayOrderStatusResp, error) {
	client := pb.NewPayServerClient(m.cli.Conn())
	return client.PayOrderStatus(ctx, in, opts...)
}

func (m *defaultPayServer) PayApprove(ctx context.Context, in *PayApproveReq, opts ...grpc.CallOption) (*PayApproveResp, error) {
	client := pb.NewPayServerClient(m.cli.Conn())
	return client.PayApprove(ctx, in, opts...)
}
