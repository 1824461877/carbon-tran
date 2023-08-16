package exchange

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
	hu "hub/utils"
	paypb "pay/pb"
	"time"
	tradedb "trade/pb"
)

type AssetExchangeTransactionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetExchangeTransactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetExchangeTransactionLogic {
	return &AssetExchangeTransactionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetExchangeTransactionLogic) AssetExchangeTransaction(req *types.ExchangeAssetTransactionReq) (resp *types.ExchangeAssetTransactionResp, err error) {
	// todo: add your logic here and delete this line
	var (
		//val        string
		TradeToken string
		PayToken   string
		times      = time.Now()
		as         = &model.AssetsSell{}
		payResp    *paypb.PayExecutionResp
		//_          *tradedb.TradeExecutionResp
	)

	if as, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByExId(l.ctx, req.ExId); err != nil {
		return nil, err
	}
	//if val, err = l.svcCtx.Redis.GetCtx(l.ctx, "exchange_"+req.ExId); err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//
	//if err != json.Unmarshal([]byte(val), &as) {
	//	return nil, err
	//}

	if (as.Number - req.Number) < 0 {
		return nil, errors.New("wrong number of transactions")
	}

	if TradeToken, err = l.TradeExecution(types.TradeOrder{
		CarbonAssetId: as.AssId,
		Initiator:     l.ctx.Value("uid").(string),
		Recipient:     as.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(l.svcCtx.Config.ServiceJwtSign.TradeServiceAuth.JwtSignExpire) * time.Minute).Unix(),
		},
	}); err != nil {
		return nil, err
	}

	if PayToken, err = l.PayExecution(types.PayOrder{
		InitiatorWalletId: req.InitiatorWalletId,
		RecipientWalletId: as.CollectionWalletId,
		PayAmount:         as.Amount * float64(req.Number),
		Initiator:         l.ctx.Value("uid").(string),
		Recipient:         as.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(l.svcCtx.Config.ServiceJwtSign.PayServiceAuth.JwtSignExpire) * time.Minute).Unix(),
		},
	}); err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.ServiceRpc.TradeRpc.TradeExecution(l.ctx, &tradedb.TradeReq{
		TaskToken: TradeToken,
		ReqTime:   times.String(),
	}); err != nil {
		return nil, err
	}

	if payResp, err = l.svcCtx.ServiceRpc.PayRpc.PayExecution(l.ctx, &paypb.PayReq{
		TaskToken: PayToken,
		ReqTime:   times.String(),
	}); err != nil {
		return nil, err
	}

	if payResp.PayStatus == types.Paid {
		var (
			assets *model.Assets
			errs   error
		)
		assets, errs = l.svcCtx.MysqlServiceContext.Assets.FindAssIdOne(l.ctx, as.AssId)
		if err != nil {
			return nil, errs
		}

		if (assets.Number - req.Number) < 0 {
			return nil, errors.New("wrong number of transactions")
		}

		var (
			VersTail int64
			OldAssID string
		)
		OldAssID = assets.AssId
		VersTail = assets.VersTail
		assets.Number = assets.Number - req.Number
		assets.VersTail = assets.VersTail - req.Number
		if assets.Number == 0 {
			if errs = l.svcCtx.MysqlServiceContext.Assets.Delete(l.ctx, assets.Id); err != nil {
				return nil, errs
			}
		} else {
			if errs = l.svcCtx.MysqlServiceContext.Assets.Update(l.ctx, assets); err != nil {
				return nil, errs
			}
		}
		assets.AssId = hu.AID()
		assets.UserId = l.ctx.Value("uid").(string)
		assets.Number = req.Number
		assets.VersHead = assets.VersTail + 1
		assets.VersTail = VersTail
		assets.Hid = fmt.Sprintf("%v-%v-%v", hu.HID(), assets.VersHead, assets.VersTail)
		assets.CreateTime = time.Now()
		assets.Listing = false
		if _, errs = l.svcCtx.MysqlServiceContext.Assets.Insert(l.ctx, assets); errs != nil {
			return nil, errs
		}

		if (as.Number - req.Number) == 0 {
			_, err = l.svcCtx.Redis.DelCtx(l.ctx, "exchange_"+req.ExId)
			if err != nil {
				return nil, err
			}
			_ = l.svcCtx.MysqlServiceContext.Assets.UpdateListing(l.ctx, &model.Assets{
				AssId:   OldAssID,
				Listing: false,
			})
		} else {
			as.Number = as.Number - req.Number
			if err = l.svcCtx.MysqlServiceContext.AssetsSell.Update(l.ctx, as); err != nil {
				return nil, err
			}
		}
	} else {
		return nil, errors.New("pay error")
	}

	return &types.ExchangeAssetTransactionResp{
		Code:    types.SuccessCode,
		Messing: "transaction success",
	}, nil
}

func (l *AssetExchangeTransactionLogic) TradeExecution(ro types.TradeOrder) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ro)
	return token.SignedString([]byte(l.svcCtx.Config.ServiceJwtSign.TradeServiceAuth.JwtSignKey))
}

func (l *AssetExchangeTransactionLogic) PayExecution(po types.PayOrder) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, po)
	return token.SignedString([]byte(l.svcCtx.Config.ServiceJwtSign.PayServiceAuth.JwtSignKey))
}
