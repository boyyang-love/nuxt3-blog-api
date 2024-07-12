package minio

import (
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"context"
	"github.com/minio/minio-go/v7"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBucketLogic {
	return &CreateBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBucketLogic) CreateBucket(req *types.CreateBucketReq) (resp *types.CreateBucketRes, err error) {
	exists, err := l.svcCtx.MinIoClient.BucketExists(l.ctx, req.Name)
	if err != nil {
		return nil, err
	}

	if exists {
		return &types.CreateBucketRes{Base: types.Base{
			Code: 1,
			Msg:  "储存桶已存在",
		}}, nil
	}

	if err = l.svcCtx.
		MinIoClient.
		MakeBucket(
			l.ctx,
			req.Name,
			minio.MakeBucketOptions{Region: "zh-east-1"},
		); err != nil {
		return nil, err
	}

	return &types.CreateBucketRes{
		Base: types.Base{
			Code: 1,
			Msg:  "储存桶创建成功",
		},
	}, nil
}
