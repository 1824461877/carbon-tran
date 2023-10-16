package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"trade/internal/svc"
	"trade/internal/types"
	"trade/model"
	"trade/pb"
	"trade/utils"
)

type TradeOrderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTradeOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TradeOrderUpdateLogic {
	return &TradeOrderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TradeOrderUpdateLogic) TradeOrderUpdate(in *pb.TradeOrderUpdateReq) (*pb.TradeOrderUpdateResp, error) {
	var (
		err error
		//	data *model.TradeOrder
	)

	var (
		assets     *model.Assets
		tradeOrder *model.TradeOrder
		errs       error
	)

	if tradeOrder, err = l.svcCtx.MysqlServiceContext.TradeOrder.FindOneByPayOrderId(l.ctx, in.PayOrderId); err != nil {
		return nil, err
	}

	if in.UpdateState == types.Expire {
		if tradeOrder.TradeStatus == types.Completed || tradeOrder.TradeStatus == types.Expire {
			return &pb.TradeOrderUpdateResp{
				Message: "dont't update state",
			}, nil
		}
		if err = l.svcCtx.MysqlServiceContext.TradeOrder.UpdateTradeStatus(l.ctx, in.PayOrderId, types.Expire); err != nil {
			return nil, err
		}
		return &pb.TradeOrderUpdateResp{
			Message: fmt.Sprintf("update state: %v", types.Expire),
		}, nil
	} else {
		if err = l.svcCtx.MysqlServiceContext.TradeOrder.UpdateTradeStatus(l.ctx, in.PayOrderId, in.UpdateState); err != nil {
			return nil, err
		}
	}

	assets, errs = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, tradeOrder.CarbonAssetId)
	if err != nil {
		return nil, errs
	}

	if (assets.Number - tradeOrder.Number) < 0 {
		return nil, errors.New("wrong number of transactions")
	}
	var (
		VersTail int64
		OldAssID string
	)
	OldAssID = assets.AssId
	VersTail = assets.VersTail
	assets.Number = assets.Number - tradeOrder.Number
	assets.VersTail = assets.VersTail - tradeOrder.Number

	if assets.Number == 0 {
		if errs = l.svcCtx.MysqlServiceContext.Assets.Delete(l.ctx, assets.Id); errs != nil {
			return nil, errs
		}
	} else {
		if errs = l.svcCtx.MysqlServiceContext.Assets.Update(l.ctx, assets); errs != nil {
			return nil, errs
		}
	}

	assets.AssId = utils.AID()
	assets.UserId = tradeOrder.Initiator
	assets.Number = tradeOrder.Number
	assets.VersHead = assets.VersTail + 1
	assets.VersTail = VersTail
	assets.Hid = fmt.Sprintf("%v-%v-%v", utils.HID(), assets.VersHead, assets.VersTail)
	assets.CreateTime = time.Now()
	assets.Listing = false
	assets.RetireNumber = 0
	if _, errs = l.svcCtx.MysqlServiceContext.Assets.Insert(l.ctx, assets); errs != nil {
		return nil, errs
	}

	var (
		as = &model.AssetsSell{}
	)
	as, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByExId(l.ctx, tradeOrder.ExchangeAssetID)
	if err != nil {
		return nil, err
	}
	if (as.Number - tradeOrder.Number) <= 0 {
		//_, err = l.svcCtx.Redis.DelCtx(l.ctx, "exchange_"+req.ExId)
		//if err != nil {
		//	return nil, err
		//}
		fmt.Println(as.Number, tradeOrder.Number, "---0000000000000000000000000")
		if err = l.svcCtx.MysqlServiceContext.AssetsSell.DeleteExID(l.ctx, tradeOrder.ExchangeAssetID); err != nil {
			return nil, err
		}
		_ = l.svcCtx.MysqlServiceContext.Assets.UpdateListing(l.ctx, &model.Assets{
			AssId:   OldAssID,
			Listing: false,
		})
	} else {
		as.Number = tradeOrder.Number
		as.ExId = tradeOrder.ExchangeAssetID
		if as.Number <= 0 {
			if err = l.svcCtx.MysqlServiceContext.AssetsSell.DeleteExID(l.ctx, as.ExId); err != nil {
				return nil, err
			}
		} else {
			if err = l.svcCtx.MysqlServiceContext.AssetsSell.UpdateNumber(l.ctx, as); err != nil {
				return nil, err
			}
		}
	}

	return &pb.TradeOrderUpdateResp{
		Message: "successful transaction",
	}, nil
}
