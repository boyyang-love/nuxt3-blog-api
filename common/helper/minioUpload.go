package helper

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

type MinioFileUploadParams struct {
	Ctx         context.Context
	MinioClient *minio.Client
	Buf         *bytes.Buffer
	Path        string
	Filename    string
}

func MinioFileUpload(params *MinioFileUploadParams) (url string, err error) {

	cloudPath := fmt.Sprintf("%s/%s", params.Path, params.Filename)

	reader := bytes.NewReader(params.Buf.Bytes())

	_, err = params.MinioClient.PutObject(
		params.Ctx,
		"boyyang",
		cloudPath,
		reader,
		reader.Size(),
		minio.PutObjectOptions{},
	)

	if err == nil {
		return cloudPath, nil
	} else {
		return "", err
	}
}
