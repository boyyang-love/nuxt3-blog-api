package blog

import (
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlogListLogic {
	return &BlogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlogListLogic) BlogList(req *types.BlogListReq) (resp *types.BlogListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
