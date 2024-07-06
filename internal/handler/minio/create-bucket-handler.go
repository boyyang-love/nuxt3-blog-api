package minio

import (
	"net/http"

	"blog_backend/internal/logic/minio"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateBucketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBucketReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := minio.NewCreateBucketLogic(r.Context(), svcCtx)
		resp, err := l.CreateBucket(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
