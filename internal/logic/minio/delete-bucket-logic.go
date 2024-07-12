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

	if err = l.svcCtx.
		MinIoClient.
		RemoveBucket(l.ctx, req.Name); err != nil {
		return nil, err
	}

	return &types.DeleteBucketRes{
		Base: types.Base{
			Code: 1,
			Msg:  "删除成功",
		},
	}, nil
}
