package tag

import (
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTagByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagByUserIdLogic {
	return &ListTagByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTagByUserIdLogic) ListTagByUserId(req *types.ListTagUserIdReq) (resp *types.ListTagUserIdRes, err error) {

	var tags []types.TagInfo
	var t []models.Tag
	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Order("created desc").
		Preload("Article", func(db *gorm.DB) *gorm.DB {
			return db.Select("id")
		}).
		Select("id", "uid", "tag_name", "type", "user_id").
		Where("user_id = ? and type = ?", req.UserId, req.Type).
		Find(&t).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&tags, &t)

	return &types.ListTagUserIdRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.ListTagUserIdResData{Tags: tags},
	}, nil
}
