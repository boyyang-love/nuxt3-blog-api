package upload

import (
	"blog_backend/common/errorx"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDeleteLogic) FileDelete(req *types.FileDeleteReq) (resp *types.FileDeleteRes, err error) {
	if err = l.del(req.Id, req.FilePath); err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	return
}

func (l *FileDeleteLogic) del(id uint, path string) (err error) {
	var uploads []models.Upload
	l.svcCtx.DB.
		Model(&models.Upload{}).
		Where("id != ? and file_path = ?", id, path).
		Find(&uploads)

	// 如何只有自己一个，则删除数据库记录
	if len(uploads) == 0 {
		if err = l.svcCtx.DB.
			Transaction(func(tx *gorm.DB) error {
				if err = l.delCloudDb(path); err != nil {
					return err
				}

				if err = l.delDb(id); err != nil {
					return err
				}
				return nil
			}); err != nil {
			return err
		}

		return nil
	}

	// 仅仅删除数据库数据 不删除对象存储数据
	if err = l.delDb(id); err != nil {
		return err
	}

	return nil
}

// 删除数据库数据
func (l *FileDeleteLogic) delDb(id uint) error {
	if err := l.svcCtx.DB.
		Model(&models.Upload{}).
		Select("id").
		Where("id = ?", id).
		Delete(&models.Upload{
			Id: id,
		}).
		Error; err != nil {
		return err
	}

	return nil
}

// 删除对象存储数据
func (l *FileDeleteLogic) delCloudDb(path string) error {
	fmt.Println(path, "hhhhhhhh")
	if _, err := l.svcCtx.Client.
		Object.
		Delete(
			context.Background(),
			path,
		); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
