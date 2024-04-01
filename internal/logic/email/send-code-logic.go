package email

import (
	"blog_backend/common/helper"
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

	err = helper.SendEmail(req.Email, fmt.Sprintf("您的验证码是：%s", code))
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Cache.Set(req.Email, []byte(code))
	if err != nil {
		return nil, err
	}

	return &types.EmailSendCodeRes{
		Message: "验证码发送成功",
	}, nil
}
