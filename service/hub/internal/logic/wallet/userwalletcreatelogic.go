package wallet

import (
	"context"
	"errors"
	"hub/model"
	uitls "hub/utils"
	"strconv"
	"time"

	"hub/internal/svc"
	"hub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWalletCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWalletCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWalletCreateLogic {
	return &UserWalletCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWalletCreateLogic) UserWalletCreate(req *types.UserWalletCreateReq) (resp *types.UserWalletCreateResp, err error) {
	// todo: add your logic here and delete this line
	var (
		times    time.Time
		wid      string
		salt     string
		errs     error
		password string
	)

	if uitls.CheckPassword(req.Password) {
		return nil, errors.New("password security is too low")
	}

	times = time.Now()
	wid = uitls.WID()
	salt = uitls.WID() + strconv.FormatInt(times.UnixNano(), 10)
	password = uitls.GenSaltPassword(salt, req.Password)
	if _, errs = l.svcCtx.MysqlServiceContext.UserWallet.Insert(l.ctx, &model.UserWallet{
		WalletId:   wid,
		Name:       req.Name,
		UserId:     l.ctx.Value("uid").(string),
		Salt:       salt,
		Password:   password,
		CreateTime: times,
		UpdateTime: times,
	}); errs != nil {
		return nil, errs
	}

	return &types.UserWalletCreateResp{
		Code:    types.SuccessCode,
		Message: "Successfully created wallet",
	}, errs
}
