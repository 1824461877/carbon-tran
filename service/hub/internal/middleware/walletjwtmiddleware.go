package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/config"
	"hub/internal/types"
	"hub/model"
	"net/http"
)

type WalletJwtMiddleware struct {
	wallet         model.UserWalletModel
	HubServiceAuth config.ServiceAuth
}

func NewWalletJwtMiddleware(HubServiceAuth config.ServiceAuth, wallet model.UserWalletModel) *WalletJwtMiddleware {
	return &WalletJwtMiddleware{
		wallet:         wallet,
		HubServiceAuth: HubServiceAuth,
	}
}

func (m *WalletJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			token string
			err   error
			aj    *AuthJwt
		)

		token = r.Header.Get("Authorization")
		if token == "" && len(token) <= 0 {
			httpx.ErrorCtx(r.Context(), w, TokenExpired)
			return
		}

		j := NewJwt(m.HubServiceAuth.JwtSignKey)
		if aj, err = j.ParseToken(token); err != nil {
			if err == TokenExpired {
				httpx.ErrorCtx(r.Context(), w, TokenExpired)
				return
			}
			httpx.ErrorCtx(r.Context(), w, TokenInvalid)
			return
		}

		var uw *model.UserWallet
		if uw, err = m.wallet.FindOneByWalletId(context.Background(), aj.Id); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Err{
				Code:    types.InternalErrorCode,
				Err:     errors.New("wallet not found").Error(),
				Message: "",
			})
			return
		} else {
			//if uw.Password != aj.Password {
			//	httpx.OkJsonCtx(r.Context(), w, types.Err{
			//		Code:    types.ValidErrorCode,
			//		Err:     TokenInvalid.Error(),
			//		Message: "",
			//	})
			//	return
			//}
		}
		fmt.Println(uw)

		r.Header.Set("wid", aj.Id)
		next(w, r)
	}
}
