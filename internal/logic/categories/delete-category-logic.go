package categories

import (
	"blog_backend/models"
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.DeleteCategorieReq) (resp *types.DetleteCategorieRes, err error) {

	if err = l.svcCtx.DB.
		Model(&models.Categories{}).
		Where("id = ?", req.Id).
		Delete(&models.Categories{}).
		Error; err != nil {
		return nil, err
	}

	return &types.DetleteCategorieRes{Base: types.Base{
		Code: 1,
		Msg:  "删除成功！",
	}}, nil
}
