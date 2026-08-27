package link

import (
	"blog_backend/models"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLinkLogic {
	return &CreateLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLinkLogic) CreateLink(req *types.LinkCreateReq) (resp *types.LinkCreateRes, err error) {
	// 从上下文获取用户UID
	uidValue := l.ctx.Value("Uid")
	if uidValue == nil {
		logx.Error("用户未登录，ctx中没有uid")
		return nil, errors.New("请先登录后再申请友链")
	}

	userUid, ok := uidValue.(string)
	if !ok || userUid == "" {
		logx.Errorf("用户UID类型断言失败，值: %v", uidValue)
		return nil, errors.New("用户信息异常")
	}

	logx.Infof("用户 %s 申请友链", userUid)

	// 检查该用户是否已申请过友链
	isExist, err := l.getLinkByUserUid(userUid)
	if err != nil {
		logx.Errorf("检查用户友链失败: %v", err)
		return nil, err
	}

	if isExist {
		return nil, errors.New("您已提交过友链申请，请等待审核")
	}

	// 创建友链
	link := &models.Links{
		WebsiteName: req.WebsiteName,
		WebsiteUrl:  req.WebsiteUrl,
		WebsiteIcon: req.WebsiteIcon,
		WebsiteDesc: req.WebsiteDesc,
		UserUid:     userUid,
	}

	if err := l.svcCtx.DB.Model(&models.Links{}).Create(link).Error; err != nil {
		logx.Errorf("创建友链失败: %v", err)
		return nil, fmt.Errorf("创建友链失败: %v", err)
	}

	logx.Infof("用户 %s 友链申请成功", userUid)

	return &types.LinkCreateRes{
		Base: types.Base{
			Code: 1,
			Msg:  "创建成功",
		},
	}, nil
}

func (l *CreateLinkLogic) getLinkByUserUid(userUid string) (isExist bool, err error) {
	var link models.Links
	if err := l.svcCtx.
		DB.
		Model(&models.Links{}).
		Where("user_uid = ?", userUid).
		First(&link).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}
