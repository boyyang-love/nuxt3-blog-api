package helper

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
)

func ImageWH(fileHeader *multipart.FileHeader) (w int, h int, err error) {

	f, err := fileHeader.Open()
	if err != nil {
		return 0, 0, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return 0, 0, err
	}

	g := img.Bounds()

	fmt.Println(g.Dx(), g.Dy())
	return g.Dx(), g.Dy(), nil
}
