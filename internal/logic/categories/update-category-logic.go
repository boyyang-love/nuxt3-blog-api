package categories

import (
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategorieReq) (resp *types.UpdateCategorieRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	if err = l.svcCtx.DB.
		Model(&models.Categories{}).
		Select("name", "cover", "des").
		Where("id = ? and user_id = ?", req.Id, userid).
		Updates(&models.Categories{
			Name:  req.Name,
			Cover: req.Cover,
			Des:   req.Des,
		}).
		Error; err != nil {
		return nil, err
	}
	return &types.UpdateCategorieRes{
		Base: types.Base{
			Code: 1,
			Msg:  "修改分类成功！",
		}}, nil
}
