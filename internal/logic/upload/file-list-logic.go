package upload

import (
	"blog_backend/models"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListLogic) FileList(req *types.FileListReq) (resp *types.FileListRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	var count int64
	var fileInfo []models.Upload
	var infos []types.FileInfo

	print(userid, req.Type, req.Limit)

	DB := l.svcCtx.DB.
		Model(&models.Upload{}).
		Order("Updated desc").
		Order("Created desc").
		Select("id", "file_name", "file_path", "origin_file_path", "w", "h", "status").
		Where("user_id = ? and type = ?", userid, req.Type)

	if req.Page == 0 || req.Limit == 0 {
		if err = DB.
			Find(&fileInfo).
			Count(&count).
			Error; err != nil {
			return nil, err
		}
	} else {
		if err = DB.
			Offset((req.Page - 1) * req.Limit).
			Limit(req.Limit).
			Find(&fileInfo).
			Offset(-1).
			Count(&count).
			Error; err != nil {
			return nil, err
		}
	}

	_ = copier.Copy(&infos, &fileInfo)

	return &types.FileListRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.FileListResdata{
			Count: count,
			Infos: infos,
		},
	}, nil
}
