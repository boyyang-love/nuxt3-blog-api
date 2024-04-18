package blog

import (
	"blog_backend/models"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBlogLogic {
	return &UpdateBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBlogLogic) UpdateBlog(req *types.UpdateBlogReq) (resp *types.UpdateBlogRes, err error) {
	userId, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	tags, err := l.getTag(req.Tags)
	if err != nil {
		return nil, err
	}

	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Select("title", "des", "cover", "content", "keywords", "categories_id").
		Where("id= ? and user_id = ?", req.Id, userId).
		Updates(&models.Article{
			Title:        req.Title,
			Des:          req.Des,
			Cover:        req.Cover,
			Content:      req.Content,
			Keywords:     req.Keywords,
			CategoriesId: req.CategoryId,
		}).
		Error; err != nil {
		return nil, err
	}

	if err = l.svcCtx.DB.
		Model(&models.Article{Id: req.Id, UserId: uint(userId)}).
		Association("Tag").
		Replace(tags); err != nil {
		return nil, err
	}

	return &types.UpdateBlogRes{
		Base: types.Base{
			Code: 1,
			Msg:  "更新成功",
		},
	}, nil
}

func (l *UpdateBlogLogic) getTag(ids []uint) (tags []*models.Tag, err error) {
	if len(ids) == 0 {
		return tags, nil
	}

	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Select("id", "type").
		Where("type = ?", "article").
		Find(&tags, ids).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return tags, err
	}

	return tags, err
}
