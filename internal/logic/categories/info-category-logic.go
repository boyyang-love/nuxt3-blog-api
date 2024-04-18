package categories

import (
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type InfoCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoCategoryLogic {
	return &InfoCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoCategoryLogic) InfoCategory(req *types.InfoCategorieReq) (resp *types.InfoCategorieRes, err error) {
	var category []models.Categories
	var infos []types.CategorieInfo
	if err = l.svcCtx.DB.
		Order("created desc").
		Model(&models.Categories{}).
		Where("user_id = ?", req.UserId).
		Select("id", "name", "cover", "des").
		Find(&category).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&infos, category)

	return &types.InfoCategorieRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.InfoCategorieResData{
			Info: infos,
		},
	}, nil
}
