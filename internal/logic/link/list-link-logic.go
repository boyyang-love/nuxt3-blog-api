package link

import (
	"blog_backend/models"
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
	var links []types.LinkListItem
	var count int64
	if err = l.svcCtx.DB.
		Model(&models.Links{}).
		Find(&links).
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	return &types.LinkListRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.LinkListData{
			List: links,
		},
		Count: count,
	}, nil
}
