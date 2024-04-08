package blog

import (
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchBlogByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchBlogByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchBlogByIdsLogic {
	return &SearchBlogByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchBlogByIdsLogic) SearchBlogByIds(req *types.BlogSearchByIdsReq) (resp *types.BlogSearchByIdsRes, err error) {
	var article []models.Article
	var list []types.BlogSearchByIdsListInfo
	if err = l.svcCtx.DB.
		Order("updated desc").
		Model(&models.Article{}).
		Select("id", "uid", "title", "des", "cover").
		Where("id in ?", req.Ids).
		Find(&article).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&list, &article)

	return &types.BlogSearchByIdsRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.BlogSearchByIdsResData{
			Info: list,
		},
	}, nil
}
