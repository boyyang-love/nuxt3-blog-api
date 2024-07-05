package link

import (
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStatusLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLinkLogic {
	return &UpdateStatusLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStatusLinkLogic) UpdateStatusLink(req *types.LinkStatusUpdateReq) (resp *types.LinkStatusUpdateRes, err error) {
	// todo: add your logic here and delete this line

	return
}
