package exchange

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/exchange"
	"hub/internal/svc"
	"hub/internal/types"
)

func GetExchangeAssetDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExchangeAssetDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := exchange.NewGetExchangeAssetDetailsLogic(r.Context(), svcCtx)
		resp, err := l.GetExchangeAssetDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
