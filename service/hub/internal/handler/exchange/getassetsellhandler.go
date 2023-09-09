package exchange

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/exchange"
	"hub/internal/svc"
	"hub/internal/types"
	"net/http"
)

func GetAssetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAssetSellReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if r.Header.Get("uid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("uid"))
		l := exchange.NewGetAssetLogic(ctx, svcCtx)
		resp, err := l.GetAsset(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
