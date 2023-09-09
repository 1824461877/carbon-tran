package personal

import (
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/personal"
	"hub/internal/svc"
	"hub/internal/types"
)

func PayOrderPaymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if r.Header.Get("uid") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
			return
		}
		l := personal.NewPayOrderPaymentLogic(r.Context(), svcCtx)
		resp, err := l.PayOrderPayment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
