package personal

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"hub/internal/logic"
	"hub/internal/svc"
	"hub/internal/types"
	"net/http"
)

func GetRetireCertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RetireFile
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewRetireFile(r.Context(), svcCtx)
		file, err := l.RetireFile("../../upload_file/" + req.RID + ".png")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			_, _ = w.Write(file)
		}
	}
}
