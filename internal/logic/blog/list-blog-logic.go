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

	DB := l.svcCtx.DB.Model(&models.Article{})

	if req.Type == "top" {
		DB = DB.Order("star desc").Order("created desc")
	}

	if req.Type == "recently" {
		DB = DB.Order("updated desc")
	}

	if req.Type == "created" {
		DB = DB.Order("created desc")
	}

	if err = DB.
		Preload("User").
		Preload("Tag").
		Preload("Comment").
		Preload("Comment.User").
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
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.ListBlogResData{
			Count: count,
			List:  lists,
		},
	}, nil
}
