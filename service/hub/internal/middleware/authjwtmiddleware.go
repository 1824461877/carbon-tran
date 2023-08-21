package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/config"
	"hub/internal/types"
	"net/http"
)

type AuthJwtMiddleware struct {
	DB          sqlx.SqlConn
	UserService config.ServiceAuth
}

func NewAuthJwtMiddleware(UserService config.ServiceAuth, userDB sqlx.SqlConn) *AuthJwtMiddleware {
	return &AuthJwtMiddleware{
		DB:          userDB,
		UserService: UserService,
	}
}

func (m *AuthJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			token string
			err   error
			aj    *AuthJwt
		)
		token = r.Header.Get("Authorization")
		if token == "" && len(token) <= 0 {
			httpx.OkJsonCtx(r.Context(), w, types.Err{
				Code:    types.TokenErrorCode,
				Err:     TokenExpired.Error(),
				Message: "",
			})
			return
		}
		j := NewJwt(m.UserService.JwtSignKey)
		if aj, err = j.ParseToken(token); err != nil {
			if err == TokenExpired {
				httpx.OkJsonCtx(r.Context(), w, types.Err{
					Code:    types.TokenErrorCode,
					Err:     TokenExpired.Error(),
					Message: "",
				})
				return
			}
			httpx.ErrorCtx(r.Context(), w, TokenInvalid)
			return
		}
		var out = struct {
			Password string `db:"password"`
		}{}
		query := fmt.Sprintf("select * from %s where `uid` = ?", "rwen_admin")
		if err = m.DB.QueryRowCtx(context.Background(), &out, query, aj.Id); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Err{
				Code:    types.InternalErrorCode,
				Err:     err.Error(),
				Message: "",
			})
			return
		} else {
			if out.Password != aj.Password {
				httpx.OkJsonCtx(r.Context(), w, types.Err{
					Code:    types.ValidErrorCode,
					Err:     TokenInvalid.Error(),
					Message: "",
				})
				return
			}
		}
		r.Header.Set("uid", aj.Id)
		next(w, r)
	}
}
