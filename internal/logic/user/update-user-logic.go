package user

import (
	"blog_backend/common/errorx"
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	var user []models.User
	l.svcCtx.DB.
		Model(&models.User{}).
		Select("username", "id").
		Where("username = ? and id != ?", req.Username, userid).
		Find(&user)

	if len(user) != 0 {
		return nil, errorx.NewDefaultError("用户名已存在")
	}

	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Where("id = ?", userid).
		Updates(&models.User{
			Username: req.Username,
			Avatar:   req.Avatar,
			Motto:    req.Motto,
			Tel:      req.Tel,
			Address:  req.Address,
			QQ:       req.QQ,
			Wechat:   req.Wechat,
			GitHub:   req.GitHub,
			Cover:    req.Cover,
		}).
		Error; err != nil {
		return nil, err
	}
	return &types.UpdateUserRes{Message: "信息修改成功"}, nil
}
