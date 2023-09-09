package wallet

import (
	"context"
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/wallet"
	"hub/internal/svc"
	"hub/internal/types"
)

func UserWalletAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserWalletAddReq
		if r.Header.Get("uid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
			return
		}

		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("uid"))
		l := wallet.NewUserWalletAddLogic(ctx, svcCtx)
		resp, err := l.UserWalletAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}

	}
}
