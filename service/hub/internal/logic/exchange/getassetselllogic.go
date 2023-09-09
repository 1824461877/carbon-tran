package exchange

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
)

type GetAssetSellLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAssetSellLogic {
	return &GetAssetSellLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAssetSellLogic) GetAsset(req *types.GetAssetSellReq) (resp *types.GetAssetTradeInfoResp, err error) {
	var (
		as  *model.Assets
		ass *model.AssetsSell
	)

	ass, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByExId(l.ctx, req.ExID)
	if err != nil {
		return nil, err
	}

	as, err = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, ass.AssId)
	if err != nil {
		return nil, err
	}

	return &types.GetAssetTradeInfoResp{
		Code: types.SuccessCode,
		Assets: &types.GetAssetTradeInfoOnce{
			ExId:    req.ExID,
			AID:     as.AssId,
			UserId:  as.UserId,
			Project: as.Project,
			Amount:  ass.Amount,
		},
	}, nil
}
