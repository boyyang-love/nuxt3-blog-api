package upload

import (
	"blog_backend/models"
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileStatusLogic {
	return &FileStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileStatusLogic) FileStatus(req *types.FileStatusReq) (resp *types.FileStatusRes, err error) {
	if err = l.svcCtx.DB.
		Debug().
		Model(&models.Upload{}).
		Where("id = ?", req.Id).
		Select("id", "status").
		Update("status", req.Status).
		Error; err != nil {
		return nil, err
	}
	return &types.FileStatusRes{
		Base: types.Base{
			Code: 1,
			Msg:  "图片状态更新成功",
		},
	}, nil
}
