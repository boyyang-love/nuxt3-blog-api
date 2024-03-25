package member

import (
	"net/http"

	"blog_backend/internal/logic/member"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := member.NewSignUpLogic(r.Context(), svcCtx)
		resp, err := l.SignUp(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
