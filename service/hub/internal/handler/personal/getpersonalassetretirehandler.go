package personal

import (
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/personal"
	"hub/internal/svc"
)

func GetPersonalAssetRetireHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("uid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
			return
		}
		l := personal.NewGetPersonalAssetRetireLogic(r.Context(), svcCtx)
		resp, err := l.GetPersonalAssetRetire(r.Header.Get("uid"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
