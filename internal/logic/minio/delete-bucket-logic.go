package minio

import (
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBucketLogic {
	return &DeleteBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBucketLogic) DeleteBucket(req *types.DeleteBucketReq) (resp *types.DeleteBucketRes, err error) {
	// todo: add your logic here and delete this line

	return
}
