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
	OriBuf      *bytes.Buffer
	FileHash    string
	Path        string
	OriPath     string
	Filename    string
}

type MinioFileReturnPaths struct {
	CloudPath    string
	OriCloudPath string
}

func MinioFileUpload(params *MinioFileUploadParams) (urls *MinioFileReturnPaths, err error) {

	fileName := fmt.Sprintf("%s-%s.webp", FileNameNoExt(params.Filename), params.FileHash)

	cloudPath := fmt.Sprintf("%s/%s", params.Path, fileName)
	cloudOriPath := fmt.Sprintf("%s/%s", params.OriPath, params.Filename)

	reader := bytes.NewReader(params.Buf.Bytes())
	oriReader := bytes.NewReader(params.OriBuf.Bytes())

	_, err = params.MinioClient.PutObject(
		params.Ctx,
		"boyyang",
		cloudPath,
		reader,
		reader.Size(),
		minio.PutObjectOptions{},
	)

	_, err = params.MinioClient.PutObject(
		params.Ctx,
		"boyyang",
		cloudOriPath,
		oriReader,
		oriReader.Size(),
		minio.PutObjectOptions{},
	)

	if err == nil {
		return &MinioFileReturnPaths{
			CloudPath:    cloudPath,
			OriCloudPath: cloudOriPath,
		}, nil
	} else {
		return nil, err
	}
}
