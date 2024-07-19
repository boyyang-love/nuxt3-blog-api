package email

import (
	"blog_backend/common/helper"
	"blog_backend/models"
	"context"
	"fmt"
	"math/rand"
	"time"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.EmailSendCodeReq) (resp *types.EmailSendCodeRes, err error) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	info, err := l.UserInfo()
	if err != nil {
		return nil, err
	}

	err = helper.SendEmail(
		helper.SendEmailParams{
			To:       req.Email,
			Subject:  req.Subject,
			Code:     code,
			UserInfo: info,
		},
	)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Cache.Set(fmt.Sprintf("%s-%s", req.Email, req.Type), []byte(code))
	if err != nil {
		return nil, err
	}

	return &types.EmailSendCodeRes{
		Base: types.Base{
			Code: 1,
			Msg:  "验证码发送成功",
		},
	}, nil
}

func (l *SendCodeLogic) UserInfo() (info *models.User, err error) {
	var userInfo models.User
	if err := l.svcCtx.DB.
		Model(&models.User{}).
		Where("id = ?", 1).
		First(&userInfo).Error; err != nil {
		return nil, err
	}

	return &userInfo, nil
}
