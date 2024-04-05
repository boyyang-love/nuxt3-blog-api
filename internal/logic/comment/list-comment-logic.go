package comment

import (
	"blog_backend/models"
	"context"
	"github.com/jinzhu/copier"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentLogic {
	return &ListCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentLogic) ListComment(req *types.CommentListReq) (resp *types.CommentListRes, err error) {
	var count int64
	var info []types.CommentInfo
	var comments []models.Comment

	DB := l.svcCtx.DB.Model(&models.Comment{}).Preload("User")

	if req.Type == "article" {
		DB = DB.Where("article_id = ?", req.Id)
	}

	if req.Type == "comment" {
		DB = DB.Where("comment_id = ?", req.Id)
	}

	if req.Type == "website" {
		DB = DB.Where("website_user_id = ?", req.Id)
	}

	if err := DB.Find(&comments).
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	_ = copier.Copy(&info, &comments)

	return &types.CommentListRes{
		Base: types.Base{
			Code: 1,
			Msg:  "ok",
		},
		Data: types.CommentListResData{
			Count: count,
			Info:  info,
		},
	}, nil
}
