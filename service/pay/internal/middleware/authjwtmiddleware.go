package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc"
	"pay/internal/config"
	"pay/internal/svc"
)

type AuthJwtMiddleware struct {
	DB             sqlx.SqlConn
	PayServiceAuth config.PayServiceAuth
}

func NewAuthJwtMiddleware(ctx *svc.ServiceContext, c config.Config) *AuthJwtMiddleware {
	return &AuthJwtMiddleware{
		//DB:             ctx.UserDB,
		PayServiceAuth: c.PayServiceAuth,
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

	j := NewJwt(ajw.PayServiceAuth.UserJwtSignKey)
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

	return handler(ctx, req)
}
