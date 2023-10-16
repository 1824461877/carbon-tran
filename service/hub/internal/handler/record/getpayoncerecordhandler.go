package record

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic/record"
	"hub/internal/svc"
	"hub/internal/types"
	"net/http"
)

func GetPayOnceRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayOnceRecordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := record.NewGetPayOnceRecordLogic(r.Context(), svcCtx)
		resp, err := l.PayOnceRecord(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
