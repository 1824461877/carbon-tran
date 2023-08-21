package wallet

import (
	"context"
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/wallet"
	"hub/internal/svc"
)

func UserWalletListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("uid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("uid"))
		l := wallet.NewUserWalletListLogic(ctx, svcCtx)
		resp, err := l.UserWalletList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}

	}
}
