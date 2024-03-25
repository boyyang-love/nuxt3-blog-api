package tag

import (
	"blog_backend/models"
	"context"
	"encoding/json"

	"blog_backend/internal/svc"
	"blog_backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagLogic) DeleteTag(req *types.DeleteTagReq) (resp *types.DeleteTagRes, err error) {
	userId, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if err = l.svcCtx.DB.
		Model(&models.Tag{}).
		Where("id = ? and user_id = ?", req.Id, userId).
		Delete(&models.Tag{}).
		Error; err != nil {
		return nil, err
	}

	return &types.DeleteTagRes{
		Message: "删除成功",
	}, nil
}
