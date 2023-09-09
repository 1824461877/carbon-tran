package personal

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/personal"
	"hub/internal/svc"
	"hub/internal/types"
	"net/http"
)

func GetPersonalTransactionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetObtainTransactionReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if r.Header.Get("uid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
			return
		}
		l := personal.NewGetPersonalTransactionListLogic(r.Context(), svcCtx)
		resp, err := l.GetPersonalTransactionList(r.Header.Get("uid"), &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
