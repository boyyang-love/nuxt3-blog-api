package blog

import (
	"context"
	"fmt"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBlogLogic {
	return &ListBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBlogLogic) ListBlog(req *types.ListBlogReq) (resp *types.ListBlogRes, err error) {
	// todo: add your logic here and delete this line
	fmt.Println(req.Page, req.Limit)
	return &types.ListBlogRes{
		Name: "boyyang",
	}, err
}
