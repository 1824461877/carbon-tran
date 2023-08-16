package exchange

import (
	"context"
	"fmt"
	"hub/model"

	"hub/internal/svc"
	"hub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExchangeAssetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExchangeAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExchangeAssetListLogic {
	return &GetExchangeAssetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExchangeAssetListLogic) GetExchangeAssetList(req *types.ExchangeAssetListReq) (resp *types.ExchangeAssetListResp, err error) {
	// todo: add your logic here and delete this line
	//var (
	//	val    []string
	//	cursor uint64
	//	keys   []string
	//)

	//if req.ExId == "" {
	//	keys, cursor, err = l.svcCtx.Redis.ScanCtx(l.ctx, cursor, "exchange_*", 10)
	//	if err != nil {
	//		return nil, err
	//	}
	//	val, err = l.svcCtx.Redis.MgetCtx(l.ctx, keys...)
	//	if err != nil {
	//		return nil, err
	//	}
	//} else {
	//	var valOnce string
	//	valOnce, err = l.svcCtx.Redis.GetCtx(l.ctx, "exchange_"+req.ExId)
	//	if err != nil {
	//		return nil, err
	//	}
	//	val = append(val, valOnce)
	//}

	var (
		all *[]model.AssetsSell
	)
	all, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	var list []types.ExchangeAssetOnceResp
	for _, v := range *all {
		var (
			one *model.Assets
			//ec  model.AssetsSell
		)
		one, err = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, v.AssId)
		if err != nil {
			break
		}
		list = append(list, types.ExchangeAssetOnceResp{
			ExId:         v.ExId,
			Assid:        v.AssId,
			UserId:       v.UserId,
			Number:       v.Number,
			GS:           one.GsId,
			Serial:       fmt.Sprintf("%v-%v", one.VersHead, one.VersTail),
			Project:      one.Project,
			SerialNumber: one.SerialNumber + fmt.Sprintf("-%v-%v", one.VersHead, one.VersTail),
			Source:       one.Source,
			Day:          one.Day,
			Amount:       v.Amount,
			Country:      one.Country,
			Product:      one.Product,
			ProjectType:  one.ProjectType,
			CreateTime:   v.CreateTime.Unix(),
			EndTime:      v.EndTime.Unix(),
		})
	}

	return &types.ExchangeAssetListResp{
		ExchangeAssetList: list,
	}, nil
}
