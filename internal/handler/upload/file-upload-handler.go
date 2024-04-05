package upload

import (
	"blog_backend/common/helper"
	"blog_backend/internal/logic/upload"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"net/http"
	"path"
	"time"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		// file hash
		fileHash, err := helper.MakeFileHash(file, fileHeader)
		// 用户上传名称 以及路径（blog,image）
		fileCustomName := r.PostFormValue("file_name")
		fileCustomDir := r.PostFormValue("dir")
		// userid
		userid, err := r.Context().Value("Id").(json.Number).Int64()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		info, err := IsExists(svcCtx.DB, fileHash, userid, fileCustomDir)
		if err != nil {
			fileCloudPath, err := helper.CosFileUpload(
				svcCtx.Client,
				fileHeader,
				fmt.Sprintf("BOYYANG/%d/%s", userid, fileCustomDir),
			)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}

			//插入数据库
			uploadData := models.Upload{
				Hash:     fileHash,
				FileName: fileCustomName,
				FileSize: fileHeader.Size,
				FileType: path.Ext(fileHeader.Filename),
				FilePath: fileCloudPath,
				UserId:   uint(userid),
				Type:     fileCustomDir,
			}

			info, err = AddToMysql(svcCtx.DB, &uploadData)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}

			httpx.OkJsonCtx(r.Context(), w, types.FileUploadRes{
				Base: types.Base{
					Code: 1,
					Msg:  "ok",
				},
				Data: types.FileUploadResdata{
					FileName: info.FileName,
					Path:     info.FilePath,
				},
			})
			return
		}

		l := upload.NewFileUploadLogic(r.Context(), svcCtx)

		req.FileName = info.FileName
		req.FilePath = info.FilePath

		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

func IsExists(DB *gorm.DB, fileHash string, userid int64, dirType string) (info *models.Upload, err error) {
	if err = DB.
		Model(&models.Upload{}).
		Select("id", "hash", "user_id", "file_name", "file_path", "type").
		Where("hash = ? and user_id = ? and type = ?", fileHash, userid, dirType).
		First(&info).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("文件不存在")
	}

	DB.Model(&models.Upload{}).
		Select("Updated").
		Where("hash = ? and user_id = ?", fileHash, userid).
		Update("Updated", time.Now().UnixMilli())

	return info, nil
}

func AddToMysql(DB *gorm.DB, upload *models.Upload) (resp *models.Upload, err error) {
	if err := DB.
		Model(&models.Upload{}).
		Create(&upload).
		Error; err != nil {
		return nil, err
	}

	return upload, err
}
