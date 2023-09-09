package wallet

//func UserWalletCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var req types.UserWalletCreateReq
//		if r.Header.Get("uid") == "" {
//			httpx.ErrorCtx(r.Context(), w, errors.New("uid is empty"))
//			return
//		}
//
//		if err := httpx.ParseJsonBody(r, &req); err != nil {
//			httpx.ErrorCtx(r.Context(), w, err)
//			return
//		}
//
//		ctx := context.WithValue(r.Context(), "uid", r.Header.Get("uid"))
//		l := wallet.NewUserWalletCreateLogic(ctx, svcCtx)
//		resp, err := l.UserWalletCreate(&req)
//		if err != nil {
//			httpx.ErrorCtx(r.Context(), w, err)
//		} else {
//			httpx.OkJsonCtx(r.Context(), w, resp)
//		}
//
//	}
//}
