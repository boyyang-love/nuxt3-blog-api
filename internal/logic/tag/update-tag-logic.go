package tag

import (
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.UpdateTagReq) (resp *types.UpdateTagRes, err error) {
	userId, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Where("id = ? and user_id = ?", req.Id, userId).
		Updates(&models.Tag{
			TagName: req.Name,
		}).
		Error; err != nil {
		return nil, err
	}

	return &types.UpdateTagRes{
		Message: "更新成功",
	}, nil
}
