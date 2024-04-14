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

	counts, err := l.UserCount(req.Id)
	if err != nil {
		return nil, err
	} else {
		info.BlogCount = counts.BlogCount
		info.WallpaperCount = counts.WallpaperCount
		info.TagsCount = counts.TagsCount
	}

	return &types.InfoUserRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: info,
	}, nil
}

type CountRes struct {
	BlogCount      int64
	WallpaperCount int64
	TagsCount      int64
}

func (l *InfoUserLogic) UserCount(userId uint) (count CountRes, err error) {
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Select("id").
		Where("user_id = ?", userId).
		Count(&count.BlogCount).
		Error; err != nil {
		return count, err
	}

	if err = l.svcCtx.DB.
		Model(&models.Upload{}).
		Select("id").
		Where("user_id = ? and type = ? and status = ?", userId, "images", true).
		Count(&count.WallpaperCount).
		Error; err != nil {
		return count, err
	}

	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Select("id").
		Where("user_id = ?", userId).
		Count(&count.TagsCount).
		Error; err != nil {
		return count, err
	}

	return count, nil
}
