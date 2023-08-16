package wallet

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/wallet"
	"hub/internal/svc"
	"hub/internal/types"
)

func UserWalletLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserWalletLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := wallet.NewUserWalletLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserWalletLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
