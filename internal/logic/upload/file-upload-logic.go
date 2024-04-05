package upload

import (
	"blog_backend/internal/types"
	"context"

	"blog_backend/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq) (res *types.FileUploadRes, err error) {

	return &types.FileUploadRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.FileUploadResdata{
			FileName: req.FileName,
			Path:     req.FilePath,
		},
	}, err
}
