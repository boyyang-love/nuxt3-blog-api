package upload

import (
	"context"
	"fmt"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadMinioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadMinioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadMinioLogic {
	return &FileUploadMinioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadMinioLogic) FileUploadMinio(req *types.FileUploadReq) (resp *types.FileUploadRes, err error) {
	return &types.FileUploadRes{
		Base: types.Base{
			Code: 1,
			Msg:  fmt.Sprintf("文件[%s]上传成功", req.FileName),
		},
		Data: types.FileUploadResdata{
			FileName: req.FileName,
			Path:     req.FilePath,
		},
	}, err
}
