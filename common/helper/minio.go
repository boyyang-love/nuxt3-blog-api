package helper

import "github.com/minio/minio-go"

func Minio(endpoint string, accessKey string, secretAccessKey string, secure bool) (minioClient *minio.Client, err error) {
	minioClient, err = minio.New(
		endpoint,
		accessKey,
		secretAccessKey,
		secure,
	)

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
