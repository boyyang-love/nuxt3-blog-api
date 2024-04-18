package blog

import (
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchBlogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchBlogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchBlogByIdLogic {
	return &SearchBlogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchBlogByIdLogic) SearchBlogById(req *types.BlogSearchByIdReq) (resp *types.BlogSearchByIdRes, err error) {
	var article models.Article
	var list types.ListBlogItem
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Preload("User").
		Preload("Tag").
		Preload("Comment").
		Preload("Comment.User").
		Preload("Categories").
		Where("id = ?", req.Id).
		First(&article).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&list, &article)

	return &types.BlogSearchByIdRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.BlogSearchByIdResData{Info: list},
	}, nil
}
