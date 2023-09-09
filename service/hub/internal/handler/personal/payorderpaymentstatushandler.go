package personal

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/personal"
	"hub/internal/svc"
	"hub/internal/types"
	"net/http"
)

func PayOrderPaymentStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		l := personal.NewPayOrderPaymentStatusLogic(r.Context(), svcCtx)
		resp, err := l.PayOrderPaymentStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
