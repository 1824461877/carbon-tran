package wallet

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hub/internal/svc"
	"hub/internal/types"
	"hub/model"
	public "hub/utils"
	"time"
)

type UserWalletAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWalletAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWalletAddLogic {
	return &UserWalletAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWalletAddLogic) UserWalletAdd(req *types.UserWalletAddReq) (resp *types.UserWalletCreateResp, err error) {
	// todo: add your logic here and delete this line
	var (
		times time.Time
		code  int64
		errs  error
	)

	switch req.CType {
	case "pay":
		code = types.PayType
	}

	if code != 0 {
		if _, errs = l.svcCtx.MysqlServiceContext.UserWallet.Insert(l.ctx, &model.UserWallet{
			WalletId:          public.WID(),
			WalletType:        code,
			CID:               req.CID,
			Name:              req.Name,
			UserId:            l.ctx.Value("uid").(string),
			DefaultCollection: true,
			CreateTime:        times,
			UpdateTime:        times,
		}); errs != nil {
			return nil, errs
		}
	} else {
		return &types.UserWalletCreateResp{
			Code:    types.ValidErrorCode,
			Message: "no type",
		}, errs
	}

	return &types.UserWalletCreateResp{
		Code:    types.SuccessCode,
		Message: "Successfully created wallet",
	}, errs
}
