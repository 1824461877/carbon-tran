package exchange

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
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
	var (
		exOne *model.AssetsSell
		one   *model.Assets
	)

	exOne, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByExId(l.ctx, req.ExId)
	if err != nil {
		return nil, err
	}
	one, err = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, exOne.AssId)
	if err != nil {
		return nil, err
	}
	return &types.ExchangeAssetDetailsResp{
		ExchangeAssetOnceResp: types.ExchangeAssetOnceResp{
			ExId:         exOne.ExId,
			Assid:        exOne.AssId,
			UserId:       exOne.UserId,
			Number:       exOne.Number,
			GS:           one.GsId,
			Serial:       fmt.Sprintf("%v-%v", one.VersHead, one.VersTail),
			Project:      one.Project,
			SerialNumber: one.SerialNumber + fmt.Sprintf("-%v-%v", one.VersHead, one.VersTail),
			Source:       one.Source,
			Day:          one.Day,
			Amount:       exOne.Amount,
			Country:      one.Country,
			Product:      one.Product,
			ProjectType:  one.ProjectType,
			CreateTime:   exOne.CreateTime.Unix(),
			EndTime:      exOne.EndTime.Unix(),
		},
	}, nil
}
