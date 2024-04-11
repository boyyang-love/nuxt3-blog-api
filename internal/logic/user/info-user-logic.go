package user

import (
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoUserLogic {
	return &InfoUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoUserLogic) InfoUser(req *types.InfoUserReq) (resp *types.InfoUserRes, err error) {

	var user models.User
	var info types.InfoUserResData
	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Select("id", "username", "avatar", "cover", "motto").
		Where("id = ?", req.Id).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&info, user)

	info.BlogCount, info.WallpaperCount, err = l.UserCount(req.Id)
	if err != nil {
		return nil, err
	}

	return &types.InfoUserRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: info,
	}, nil
}

func (l *InfoUserLogic) UserCount(userId uint) (blogCount int64, wallpaperCount int64, err error) {
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Select("id").
		Where("user_id = ?", userId).
		Count(&blogCount).
		Error; err != nil {
		return blogCount, wallpaperCount, err
	}

	if err = l.svcCtx.DB.
		Model(&models.Upload{}).
		Select("id").
		Where("user_id = ? and type = ?", userId, "images").
		Count(&wallpaperCount).
		Error; err != nil {
		return blogCount, wallpaperCount, err
	}

	return blogCount, wallpaperCount, nil
}
