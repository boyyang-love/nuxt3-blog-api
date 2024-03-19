package helper

import (
	"crypto/md5"
	"fmt"
	"mime/multipart"
)

func MakeFileHash(file multipart.File, fileHeader *multipart.FileHeader) (hash string, err error) {
	h := make([]byte, fileHeader.Size)
	if _, err = file.Read(h); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%x", md5.Sum(h)), nil
	}
}
