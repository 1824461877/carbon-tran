package wallet

import (
	"context"
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/wallet"
	"hub/internal/svc"
)

func UserWalletInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("wid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("wid is empty"))
			return
		}
		ctx := context.WithValue(r.Context(), "wid", r.Header.Get("wid"))
		l := wallet.NewUserWalletInfoLogic(ctx, svcCtx)
		resp, err := l.UserWalletInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
