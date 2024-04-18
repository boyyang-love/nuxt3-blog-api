package search

import (
	"blog_backend/models"
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryIdSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryIdSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryIdSearchLogic {
	return &CategoryIdSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryIdSearchLogic) CategoryIdSearch(req *types.CategoryIdSearchReq) (resp *types.CategoryIdSearchRes, err error) {
	var articles []types.SearchResDataInfo
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Select("id", "uid", "created", "title", "des", "cover").
		Where("user_id = ? and categories_id = ?", req.UserId, req.Id).
		Find(&articles).
		Error; err != nil {
		return nil, err
	}
	return &types.CategoryIdSearchRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.CategoryIdSearchResData{
			Infos: articles,
		},
	}, nil
}
