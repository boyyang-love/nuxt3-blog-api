package user

import (
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
	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Where("id = ?", userid).
		Updates(&models.User{
			// 这里可以添加更多的字段更新
			Username: req.Username,
			Avatar:   req.Avatar,
			Motto:    req.Motto,
			Email:    req.Email,
			Tel:      req.Tel,
			Address:  req.Address,
			QQ:       req.QQ,
			Wechat:   req.Wechat,
			GitHub:   req.GitHub,
		}).
		Error; err != nil {

	}
	return &types.UpdateUserRes{Message: "信息修改成功"}, nil
}
