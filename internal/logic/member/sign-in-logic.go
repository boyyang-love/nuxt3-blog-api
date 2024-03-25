package member

import (
	"blog_backend/common/errorx"
	"blog_backend/common/helper"
	"blog_backend/models"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInRes, err error) {
	var user models.User
	passwordHash, err := helper.MakeHash(req.Password)
	if err != nil {
		passwordHash = req.Password
	}

	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Where("username = ? AND password = ?", req.Username, passwordHash).
		First(&user).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, errorx.NewDefaultError("用户名或密码错误")
	} else {
		token, err := helper.GenerateJwtToken(&helper.GenerateJwtStruct{
			Id:               user.Id,
			Uid:              user.Uid,
			Username:         user.Username,
			RegisteredClaims: jwt.RegisteredClaims{},
		},
			l.svcCtx.Config.Auth.AccessSecret,
			l.svcCtx.Config.Auth.AccessExpire)
		if err != nil {
			return resp, errorx.NewDefaultError("token生成失败")
		}
		var userInfo types.UserInfo
		err = copier.Copy(&userInfo, &user)
		if err != nil {
			return nil, err
		}

		return &types.SignInRes{
			UserInfo: userInfo,
			Token:    token,
		}, err
	}
}
