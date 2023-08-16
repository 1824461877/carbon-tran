package exchange

import (
	"context"

	"hub/internal/svc"
	"hub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExchangeAssetDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExchangeAssetDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExchangeAssetDetailsLogic {
	return &GetExchangeAssetDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExchangeAssetDetailsLogic) GetExchangeAssetDetails(req *types.ExchangeAssetDetailsReq) (resp *types.ExchangeAssetDetailsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
