package exchange

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
	"time"
	"trade/pb"
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
		TradeToken string
		//PayToken string
		times    = time.Now()
		as       = &model.AssetsSell{}
		tPayResp *pb.TradeExecutionResp
		uw       *model.UserWallet
	)

	if as, err = l.svcCtx.MysqlServiceContext.AssetsSell.FindOneByExId(l.ctx, req.ExId); err != nil {
		return nil, err
	}

	uw, err = l.svcCtx.MysqlServiceContext.UserWallet.FindOneByWalletId(l.ctx, as.CollectionWalletId)
	if err != nil {
		return nil, err
	}

	if uw != nil {

		if TradeToken, err = l.TradeExecution(types.TradeOrder{
			PayAmount:       as.Amount,
			Number:          req.Number,
			ExchangeAssetID: as.ExId,
			CollectionID:    uw.WalletId,
			Initiator:       l.ctx.Value("uid").(string),
			Recipient:       as.UserId,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(l.svcCtx.Config.ServiceJwtSign.TradeServiceAuth.JwtSignExpire) * time.Minute).Unix(),
			},
		}); err != nil {
			return nil, err
		}
	} else {
		return &types.ExchangeAssetTransactionResp{
			Code:    types.ValidErrorCode,
			Messing: "no have wallet",
		}, nil
	}

	if tPayResp, err = l.svcCtx.ServiceRpc.TradeRpc.TradeExecution(l.ctx, &pb.TradeReq{
		TaskToken: TradeToken,
		ReqTime:   times.String(),
	}); err != nil {
		return nil, err
	}

	return &types.ExchangeAssetTransactionResp{
		Code:    types.SuccessCode,
		OrderId: tPayResp.TradeOrderId,
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
