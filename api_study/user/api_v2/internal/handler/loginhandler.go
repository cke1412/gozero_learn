package handler

import (
	"net/http"
	"zero_study/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero_study/api_study/user/api_v2/internal/logic"
	"zero_study/api_study/user/api_v2/internal/svc"
	"zero_study/api_study/user/api_v2/internal/types"
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
		response.Response(w, r, resp, err)
	}
}
