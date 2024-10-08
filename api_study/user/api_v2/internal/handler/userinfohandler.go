package handler

import (
	"net/http"
	"zero_study/common/response"

	"zero_study/api_study/user/api_v2/internal/logic"
	"zero_study/api_study/user/api_v2/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, r, resp, err)
	}
}
