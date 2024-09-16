package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero_study/model_study/user/api/internal/logic"
	"zero_study/model_study/user/api/internal/svc"
	"zero_study/model_study/user/api/internal/types"

	"zero_study/common/response"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(w, r, resp, err)
	}
}
