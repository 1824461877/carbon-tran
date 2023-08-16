package gsf_back

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gsf/internal/logic/gsf_back"
	"gsf/internal/svc"
	"gsf/internal/types"
)

func ChineseTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChineseTemplateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := gsf_back.NewChineseTemplateLogic(r.Context(), svcCtx)
		resp, err := l.ChineseTemplate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
