package personal

import (
	"context"
	"encoding/json"
	"errors"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
	public "hub/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonalAssetSellLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersonalAssetSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersonalAssetSellLogic {
	return &PersonalAssetSellLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersonalAssetSellLogic) PersonalAssetSell(req *types.PersonalAssetSellReq) (resp *types.PersonalAssetSellResp, err error) {
	asset, err := l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, req.AssId)
	if err != nil {
		return nil, err
	}

	uid := l.ctx.Value("uid").(string)
	if uid != asset.UserId {
		return nil, errors.New("assets does not exist")
	}

	switch req.Status {
	case types.Sell:
		_, err = l.svcCtx.MysqlServiceContext.UserWallet.FindOneByWalletId(l.ctx, req.CollectionWalletId)
		if err != nil {
			return nil, err
		}

		if req.Number <= asset.Number {
			if req.Number > asset.Number {
				return nil, errors.New("incorrect quantity")
			}
		}

		if req.Amount <= 0 {
			return nil, errors.New("incorrect amount value")
		}

		var (
			ExId = public.EID()
			now  = time.Now()
			data = &model.AssetsSell{
				ExId:               ExId,
				UserId:             uid,
				AssId:              req.AssId,
				Country:            asset.Country,
				Amount:             req.Amount,
				CollectionWalletId: req.CollectionWalletId,
				Source:             asset.Source,
				Number:             req.Number,
				CreateTime:         now,
				EndTime:            now.Add(time.Duration(req.KeepTime) * time.Minute),
			}
			marshal []byte
		)
		if _, err = l.svcCtx.MysqlServiceContext.AssetsSell.Insert(l.ctx, data); err != nil {
			return nil, err
		}

		if marshal, err = json.Marshal(data); err != nil {
			return nil, err
		} else {
			ss := int(data.EndTime.Unix() - data.CreateTime.Unix())
			err = l.svcCtx.Redis.SetexCtx(l.ctx, "exchange_"+ExId, string(marshal), ss)
			if err != nil {
				return nil, err
			}
			err = l.svcCtx.MysqlServiceContext.Assets.UpdateListing(l.ctx, &model.Assets{
				AssId:   data.AssId,
				Listing: true,
			})
			if err != nil {
				return nil, err
			}
		}

	case types.NotSell:
		var as *model.AssetsSell
		if as, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByAssId(l.ctx, req.AssId); err != nil {
			return nil, err
		}
		_, err = l.svcCtx.Redis.DelCtx(l.ctx, "exchange_"+as.ExId)
		if err != nil {
			return nil, err
		}

		if err = l.svcCtx.MysqlServiceContext.AssetsSell.Delete(l.ctx, req.AssId); err != nil {
			return nil, err
		}

		err = l.svcCtx.MysqlServiceContext.Assets.UpdateListing(l.ctx, &model.Assets{
			AssId:   asset.AssId,
			Listing: false,
		})
		if err != nil {
			return nil, err
		}

	}

	return &types.PersonalAssetSellResp{
		Code:    types.SuccessCode,
		Message: "do is success",
	}, nil
}
