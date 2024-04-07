package tag

import (
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagLogic {
	return &ListTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTagLogic) ListTag(req *types.ListTagReq) (resp *types.ListTagRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	var tags []types.TagInfo
	var t []models.Tag
	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Order("created desc").
		Preload("Article", func(db *gorm.DB) *gorm.DB {
			return db.Select("id")
		}).
		Select("id", "uid", "tag_name", "type", "user_id").
		Where("user_id = ? and type = ?", userid, req.Type).
		Find(&t).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&tags, &t)

	return &types.ListTagRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.ListTagResData{Tags: tags},
	}, nil
}
