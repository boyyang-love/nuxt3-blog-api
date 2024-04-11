package blog

import (
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchBlogByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchBlogByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchBlogByUserIdLogic {
	return &SearchBlogByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchBlogByUserIdLogic) SearchBlogByUserId(req *types.BlogSearchByUserIdReq) (resp *types.BlogSearchByUserIdRes, err error) {
	var articles []models.Article
	var infos []types.BlogSearchByUserIdInfo
	var count int64
	if err = l.svcCtx.DB.
		Order("created desc").
		Model(&models.Article{}).
		Offset((req.Page-1)*req.Limit).
		Limit(req.Limit).
		Select("id", "title", "des", "cover", "created", "updated").
		Where("user_id = ?", req.Id).
		Find(&articles).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, err
	}
	_ = copier.Copy(&infos, &articles)
	return &types.BlogSearchByUserIdRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.BlogSearchByUserIdResData{
			Count: count,
			Infos: infos,
		},
	}, nil
}
