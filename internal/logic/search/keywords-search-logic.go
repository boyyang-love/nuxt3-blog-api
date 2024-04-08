package search

import (
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeywordsSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeywordsSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeywordsSearchLogic {
	return &KeywordsSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeywordsSearchLogic) KeywordsSearch(req *types.SearchReq) (resp *types.SearchRes, err error) {
	var blog []models.Article
	var infos []types.SearchResDataInfo

	if err = l.svcCtx.DB.
		Order("created desc").
		Model(&models.Article{}).
		Select("id", "uid", "title", "des", "cover", "created").
		Where("title LIKE ? or des LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").
		Find(&blog).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&infos, blog)

	return &types.SearchRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.SearchResData{
			Infos: infos,
		},
	}, nil
}
