package personal

import (
	"context"
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/personal"
	"hub/internal/svc"
	"hub/internal/types"
)

func PersonalAssetSellHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PersonalAssetSellReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if r.Header.Get("uid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("uid"))
		l := personal.NewPersonalAssetSellLogic(ctx, svcCtx)
		resp, err := l.PersonalAssetSell(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
