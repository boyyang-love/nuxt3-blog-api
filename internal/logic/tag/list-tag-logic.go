package tag

import (
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

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
	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Order("created desc").
		Select("id", "uid", "tag_name", "type", "user_id").
		Where("user_id = ?", userid).
		Find(&tags).
		Error; err != nil {
		return nil, err
	}

	return &types.ListTagRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.ListTagResData{Tags: tags},
	}, nil
}
