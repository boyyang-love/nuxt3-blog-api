package upload

import (
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoPublicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileInfoPublicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoPublicLogic {
	return &FileInfoPublicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileInfoPublicLogic) FileInfoPublic(req *types.FileListPublicReq) (resp *types.FileListPublicRes, err error) {
	var upload []models.Upload
	var infos []types.FileListPublicResDataInfo
	var count int64
	if err = l.svcCtx.DB.
		Debug().
		Order("Updated desc").
		Model(&models.Upload{}).
		Select("id", "file_name", "file_path", "w", "h", "status").
		Where("user_id=? and type=? and status=?", req.Id, "images", true).
		Offset((req.Page - 1) * req.Limit).
		Limit(req.Limit).
		Find(&upload).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&infos, upload)

	return &types.FileListPublicRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.FileListPublicResData{
			Count: count,
			Infos: infos,
		},
	}, nil
}
