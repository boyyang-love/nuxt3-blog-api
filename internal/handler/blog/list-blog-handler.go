package blog

import (
	"net/http"

	"blog_backend/internal/logic/blog"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListBlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListBlogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewListBlogLogic(r.Context(), svcCtx)
		resp, err := l.ListBlog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			//respx.Resp(w, resp, msg)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
