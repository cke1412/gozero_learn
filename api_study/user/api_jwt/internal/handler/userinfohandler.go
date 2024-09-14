package handler

import (
	"net/http"

	"zero_study/api_study/user/api_jwt/internal/logic"
	"zero_study/api_study/user/api_jwt/internal/svc"
	"zero_study/common/response"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(w, r, resp, err)
	}
}
