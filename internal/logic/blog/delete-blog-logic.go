package blog

import (
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/clause"
)

type DeleteBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlogLogic {
	return &DeleteBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBlogLogic) DeleteBlog(req *types.DeleteBlogReq) (resp *types.DeleteBlogRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Select(clause.Associations).
		Delete(&models.Article{
			Id:     req.Id,
			UserId: uint(userid),
		}).
		Error; err != nil {
		return nil, err
	}

	return &types.DeleteBlogRes{Message: "删除成功"}, nil
}
