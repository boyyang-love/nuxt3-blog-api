package blog

import (
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"
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
	var articles []models.Article
	var lists []types.ListBlogItem
	var count int64
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Order("Created desc").
		Preload("User").
		Preload("Tag").
		Limit(req.Limit).
		Offset((req.Page - 1) * req.Limit).
		Find(&articles).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&lists, &articles)

	return &types.ListBlogRes{
		Count: count,
		List:  lists,
	}, nil
}
