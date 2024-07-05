package link

import (
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLinkLogic {
	return &UpdateLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLinkLogic) UpdateLink(req *types.LinkUpdateReq) (resp *types.LinkUpdateRes, err error) {
	// todo: add your logic here and delete this line

	return
}
