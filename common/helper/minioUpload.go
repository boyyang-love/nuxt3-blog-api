package helper

import (
	"fmt"
	"github.com/minio/minio-go"
	"mime/multipart"
)

type MinioFileUploadParams struct {
	MinioClient   *minio.Client
	FileHeader    *multipart.FileHeader
	Path          string
	ContentLength int64
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
