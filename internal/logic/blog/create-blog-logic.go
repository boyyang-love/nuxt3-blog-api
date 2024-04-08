package blog

import (
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlogLogic {
	return &CreateBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *CreateBlogLogic) CreateBlog(req *types.CreateBlogReq) (resp *types.CreateBlogRes, err error) {
	userId, err := l.ctx.Value("Id").(json.Number).Int64() // 用户id
	if err != nil {
		return nil, err
	}

	tags, err := l.getTag(req.Tags)
	if err != nil {
		return nil, err
	}

	article := models.Article{
		Title:    req.Title,
		Des:      req.Des,
		Cover:    req.Cover,
		Content:  req.Content,
		UserId:   uint(userId),
		Keywords: req.Keywords,
		Tag:      tags,
	}

	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Create(&article).
		Error; err != nil {
		return nil, err
	}

	return &types.CreateBlogRes{
		Base: types.Base{
			Code: 1,
			Msg:  "创建成功",
		},
		Data: types.CreateBlogResData{Id: article.Id},
	}, nil
}

func (l *CreateBlogLogic) getTag(ids []uint) (tags []*models.Tag, err error) {
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
