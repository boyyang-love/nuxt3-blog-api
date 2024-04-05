package user

import (
	"blog_backend/common/helper"
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.UpdateUserPasswordReq) (resp *types.UpdateUserPasswordRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	password, err := helper.MakeHash(req.Password)
	if err != nil {
		return nil, err
	}

	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Select("id", "password").
		Where("id = ?", userid).
		Update("password", password).
		Error; err != nil {
		return nil, err
	}

	return &types.UpdateUserPasswordRes{
		Base: types.Base{
			Code: 1,
			Msg:  "更新用户密码成功",
		},
	}, nil
}
