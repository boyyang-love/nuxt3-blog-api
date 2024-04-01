package comment

import (
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CommentCreateReq) (resp *types.CommentCreateRes, err error) {
	userid, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	comment := models.BaseComment{
		Content:       req.Content,
		ArticleId:     req.ArticleId,
		CommentId:     req.CommentId,
		WebsiteUserId: req.WebsiteUserId,
		Type:          req.Type,
		UserId:        uint(userid),
	}
	if err = l.svcCtx.DB.
		Model(&models.BaseComment{}).
		Create(&comment).
		Error; err != nil {
		return nil, err
	}

	return &types.CommentCreateRes{Message: "评论成功"}, nil
}
