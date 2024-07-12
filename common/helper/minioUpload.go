package helper

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type MinioFileUploadParams struct {
	Ctx         context.Context
	MinioClient *minio.Client
	FileHeader  *multipart.FileHeader
	Path        string
}

func MinioFileUpload(params *MinioFileUploadParams) (url string, err error) {

	cloudPath := fmt.Sprintf("%s/%s", params.Path, params.FileHeader.Filename)
	f, _ := params.FileHeader.Open()
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	_, err = params.MinioClient.PutObject(
		params.Ctx,
		"boyyang",
		cloudPath,
		f,
		params.FileHeader.Size,
		minio.PutObjectOptions{},
	)

	if err == nil {
		return cloudPath, nil
	} else {
		return "", err
	}
}
