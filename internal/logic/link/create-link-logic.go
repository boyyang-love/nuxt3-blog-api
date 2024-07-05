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

	isExist, err := l.getLinkByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if isExist {
		return nil, errors.New("该邮箱已存在")
	}

	ok, err := l.codeVerify(req.Email, req.Code)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("验证码错误")
	}

	err = l.svcCtx.DB.
		Model(&models.Links{}).
		Create(&models.Links{
			WebsiteName: req.WebsiteName,
			WebsiteUrl:  req.WebsiteUrl,
			WebsiteIcon: req.WebsiteIcon,
			Email:       req.Email,
		}).Error
	if err != nil {
		return nil, err
	}

	return &types.LinkCreateRes{
		Base: types.Base{
			Code: 1,
			Msg:  "创建成功",
		},
	}, nil
}

func (l *CreateLinkLogic) getLinkByEmail(email string) (isExist bool, err error) {
	var links []models.Links
	if err := l.svcCtx.
		DB.
		Debug().
		Model(&models.Links{}).
		Select("id").
		Where("email = ?", email).
		First(&links).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		fmt.Println("1111")
		return true, nil
	}
}

func (l *CreateLinkLogic) codeVerify(email string, code string) (ok bool, err error) {
	if cacheCode, err := l.svcCtx.Cache.Get(email); err != nil {
		return false, errors.New("验证码过期")
	} else {
		if string(cacheCode) == code {
			return true, nil
		} else {
			return false, nil
		}
	}
}
