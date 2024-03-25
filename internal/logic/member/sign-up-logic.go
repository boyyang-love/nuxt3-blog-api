package member

import (
	"blog_backend/common/errorx"
	"blog_backend/common/helper"
	"blog_backend/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (resp *types.SignUpRes, err error) {
	var user models.User
	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Select("username", "email").
		Where("username = ? OR email = ?", req.Username, req.Email).
		First(&user).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		passwordHash, err := helper.MakeHash(req.Password)
		if err != nil {
			passwordHash = req.Password
		}
		if err = l.svcCtx.DB.
			Model(&models.User{}).
			Create(&models.User{
				Username: req.Username,
				Password: passwordHash,
				Email:    req.Email,
				Account:  req.Email,
			}).
			Error; err != nil {
			return resp, err
		}

		return &types.SignUpRes{Message: "注册成功"}, nil
	}

	if user.Username == req.Username {
		return resp, errorx.NewDefaultError("用户名已存在")
	}

	if user.Email == req.Email {
		return resp, errorx.NewDefaultError("邮箱已注册")
	}

	return resp, err
}
