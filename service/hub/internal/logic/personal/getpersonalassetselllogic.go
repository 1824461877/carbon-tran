package personal

import (
	"context"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonalAssetSellLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalAssetSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalAssetSellLogic {
	return &GetPersonalAssetSellLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalAssetSellLogic) GetPersonalAssetSell(req *types.GetPersonalAssetSellReq) (resp []types.GetPersonalAssetSellResp, err error) {
	// todo: add your logic here and delete this line
	var (
		list []types.GetPersonalAssetSellResp
	)
	if req.ExId == "" {
		var asList *[]model.AssetsSell
		asList, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindAllByUID(l.ctx, l.ctx.Value("uid").(string))
		if err != nil {
			return nil, err
		}

		for _, v := range *asList {
			list = append(list, types.GetPersonalAssetSellResp{
				ExId:       v.ExId,
				Assid:      v.AssId,
				UserId:     v.UserId,
				Number:     v.Number,
				CreateTime: v.CreateTime.Unix(),
				EndTime:    v.EndTime.Unix(),
			})
		}
	} else {
		var as *model.AssetsSell
		as, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByExId(l.ctx, req.ExId)
		if err != nil {
			return nil, err
		}

		list = append(list, types.GetPersonalAssetSellResp{
			ExId:       as.ExId,
			Assid:      as.AssId,
			UserId:     as.UserId,
			Number:     as.Number,
			CreateTime: as.CreateTime.Unix(),
			EndTime:    as.EndTime.Unix(),
		})
	}
	return list, nil
}
