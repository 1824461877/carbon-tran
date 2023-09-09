package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc"
	"trade/internal/config"
	"trade/internal/svc"
)

type AuthJwtMiddleware struct {
	DB             sqlx.SqlConn
	ServiceJwtSign config.ServiceJwtSign
}

func NewAuthJwtMiddleware(ctx *svc.ServiceContext, c config.Config) *AuthJwtMiddleware {
	return &AuthJwtMiddleware{
		DB:             ctx.MysqlServiceContext.UserDB,
		ServiceJwtSign: c.ServiceJwtSign,
	}
}

type TokenReq struct {
	Token string `json:"token"`
}

type TokenInterface interface {
	GetUserToken() string
}

func (ajw *AuthJwtMiddleware) Handle(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var (
		tokenInter TokenInterface
		ok         bool
		token      string
		errs       error
		aj         *AuthJwt
	)

	if tokenInter, ok = req.(TokenInterface); tokenInter == nil || !ok {
		return nil, errors.New("user_token is error")
	}

	token = tokenInter.GetUserToken()
	if token == "" && len(token) <= 0 {
		return "", errors.New("token is empty")
	}

	j := NewJwt(ajw.ServiceJwtSign.UserServiceAuth.JwtSignKey)
	if aj, errs = j.ParseToken(token); errs != nil {
		if errs == TokenExpired {
			return "", errs
		}
		return "", errs
	}

	var out = struct {
		Password string `json:"password"`
	}{}

	query := fmt.Sprintf("select * from %s where `uid` = ?", "rwen_admin")
	if errs = ajw.DB.QueryRowCtx(context.Background(), out, query, aj.Id); err != nil {
		return "", errs
	} else {
		if out.Password != aj.Password {
			return "", TokenInvalid
		}
	}

	ctx = context.WithValue(ctx, "uid", aj.Id)
	return handler(ctx, req)
}
