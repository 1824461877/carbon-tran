package wallet

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"hub/internal/middleware"
	"hub/model"
	uitls "hub/utils"
	"time"

	"hub/internal/svc"
	"hub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWalletLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWalletLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWalletLoginLogic {
	return &UserWalletLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWalletLoginLogic) UserWalletLogin(req *types.UserWalletLoginReq) (resp *types.UserWalletLoginResp, err error) {
	// todo: add your logic here and delete this line
	var (
		jwts       *middleware.JWT
		userWallet *model.UserWallet
		pass       string
		token      string
	)
	if userWallet, err = l.svcCtx.MysqlServiceContext.UserWallet.FindOneByWalletId(l.ctx, req.WalletId); err != nil {
		return nil, err
	}

	pass = uitls.GenSaltPassword(userWallet.Salt, req.Password)
	if userWallet.Password != pass {
		return nil, errors.New("password is error")
	}

	jwts = middleware.NewJwt(l.svcCtx.Config.ServiceJwtSign.HubServiceAuth.JwtSignKey)
	if token, err = jwts.CreateToken(middleware.AuthJwt{
		Id:       req.WalletId,
		Password: pass,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(l.svcCtx.Config.ServiceJwtSign.HubServiceAuth.JwtSignExpire) * time.Minute).Unix(),
		},
	}); err != nil {

		return nil, err
	}
	return &types.UserWalletLoginResp{
		Token:     token,
		LoginTime: time.Now().String(),
	}, nil
}
