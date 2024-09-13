package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero_study/api_study/user/api_v2/internal/logic"
	"zero_study/api_study/user/api_v2/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
