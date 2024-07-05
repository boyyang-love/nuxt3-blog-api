package link

import (
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLinkLogic {
	return &ListLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLinkLogic) ListLink(req *types.LinkListReq) (resp *types.LinkListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
