package categories

import (
	"net/http"

	"blog_backend/internal/logic/categories"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCategorieReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := categories.NewDeleteCategoryLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
