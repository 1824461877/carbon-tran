package wallet

import (
	"context"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWalletListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWalletListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWalletListLogic {
	return &UserWalletListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWalletListLogic) UserWalletList() (resp *types.UserWalletListResp, err error) {
	var userWalltList *[]model.UserWallet
	userWalltList, err = l.svcCtx.MysqlServiceContext.UserWallet.FindAllByUserId(l.ctx, l.ctx.Value("uid").(string))
	if err != nil {
		return nil, err
	}

	var list []types.WalletOnce
	for _, v := range *userWalltList {

		list = append(list, types.WalletOnce{
			Name:              v.Name,
			WalletId:          v.WalletId,
			DefaultCollection: v.DefaultCollection,
			WalletType:        v.WalletType,
		})
	}

	return &types.UserWalletListResp{
		WalletIdList: list,
	}, nil
}
