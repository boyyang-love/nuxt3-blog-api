package categories

import (
	"blog_backend/models"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategorieReq) (resp *types.CreateCategorieRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	isSameName, err := l.isSameName(req.Name, uint(userid))
	if err != nil {
		return nil, err
	}

	if isSameName {
		return nil, errors.New("分类已存在")
	}

	if err = l.svcCtx.DB.
		Model(&models.Categories{}).
		Select("name", "cover", "des", "user_id").
		Where("name = ? and user_id = ?", req.Name, uint(userid)).
		FirstOrCreate(&models.Categories{
			Name:   req.Name,
			Cover:  req.Cover,
			Des:    req.Des,
			UserId: uint(userid),
		}).Error; err != nil {
		return nil, err
	}
	return &types.CreateCategorieRes{
		Base: types.Base{
			Code: 1,
			Msg:  "分类创建成功",
		},
	}, nil
}

func (l *CreateCategoryLogic) isSameName(name string, userid uint) (isSame bool, err error) {
	var category models.Categories
	if err = l.svcCtx.DB.
		Model(&models.Categories{}).
		Where("name = ? and user_id = ?", name, userid).
		Select("name", "user_id").
		First(&category).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
