package upload

import (
	"blog_backend/common/helper"
	"blog_backend/internal/logic/upload"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"blog_backend/models"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func FileUploadMinioHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		// 用户上传名称 以及路径（blog,image）
		fileCustomName := r.PostFormValue("file_name")
		fileCustomDir := r.PostFormValue("dir")

		// userid
		userid, err := r.Context().Value("Id").(json.Number).Int64()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		fileHash, err := helper.MakeFileHash(file, fileHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		info, err := IsExists(svcCtx.DB, fileHash, userid, fileCustomDir)
		if err != nil {
			compressedImage, err := helper.Image2Webp(fileHeader, 80)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}

			x, y, fileSize, imgType := compressedImage.Width, compressedImage.Height, compressedImage.Size, compressedImage.Type
			fileCloudPaths, err := helper.MinioFileUpload(
				&helper.MinioFileUploadParams{
					Ctx:         r.Context(),
					MinioClient: svcCtx.MinIoClient,
					Buf:         compressedImage.Buf,
					OriBuf:      compressedImage.OriginBuf,
					FileHash:    compressedImage.Hash,
					Filename:    fileHeader.Filename,
					Path:        fmt.Sprintf("BOYYANG/%d/%s/%s", userid, fileCustomDir, "compressed"),
					OriPath:     fmt.Sprintf("BOYYANG/%d/%s/%s", userid, fileCustomDir, "origin"),
				},
			)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}

			//插入数据库
			uploadData := models.Upload{
				Hash:           fileHash,
				FileName:       helper.FileNameNoExt(fileCustomName),
				FileSize:       fileSize,
				OriginFileSize: compressedImage.OriginSize,
				FileType:       imgType,
				FilePath:       fileCloudPaths.CloudPath,
				OriginFilePath: fileCloudPaths.OriCloudPath,
				UserId:         uint(userid),
				OriginType:     compressedImage.OriginType,
				Type:           fileCustomDir,
				W:              x,
				H:              y,
			}

			info, err = AddToMysql(svcCtx.DB, &uploadData)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}

			httpx.OkJsonCtx(r.Context(), w, types.FileUploadRes{
				Base: types.Base{
					Code: 1,
					Msg:  fmt.Sprintf("文件[%s]上传成功", info.FileName),
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
