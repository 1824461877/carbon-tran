package wallet

import (
	"context"
	"hub/internal/svc"
	"hub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWalletInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWalletInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWalletInfoLogic {
	return &UserWalletInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWalletInfoLogic) UserWalletInfo() (resp *types.UserWalletInfoResp, err error) {
	// todo: add your logic here and delete this line
	userWallet, err := l.svcCtx.MysqlServiceContext.UserWallet.FindOneByWalletId(l.ctx, l.ctx.Value("wid").(string))
	if err != nil {
		return nil, err
	}

	return &types.UserWalletInfoResp{
		UserId:     userWallet.UserId,
		WalletId:   userWallet.WalletId,
		Amount:     userWallet.Amount,
		CreateTime: userWallet.CreateTime.String(),
	}, nil
}
