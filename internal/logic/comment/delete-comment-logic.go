package comment

import (
	"blog_backend/models"
	"context"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.CommentDeleteReq) (resp *types.CommentDeleteRes, err error) {

	if err = l.svcCtx.DB.
		Model(&models.Comment{}).
		Where("id = ?", req.Id).
		Delete(&models.Comment{}).
		Error; err != nil {
		return nil, err
	}

	return &types.CommentDeleteRes{
		Base: types.Base{
			Code: 1,
			Msg:  "删除评论成功",
		},
	}, nil
}
