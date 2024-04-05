package tag

import (
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTagLogic) CreateTag(req *types.CreateTagReq) (resp *types.CreateTagRes, err error) {
	userId, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Where("tag_name = ? and user_id = ?", req.Name, userId).
		FirstOrCreate(&models.Tag{
			TagName: req.Name,
			Type:    req.Type,
			UserId:  uint(userId),
		}).
		Error; err != nil {
		return nil, err
	}

	return &types.CreateTagRes{
		Base: types.Base{
			Code: 1,
			Msg:  "标签创建成功",
		},
	}, nil
}
