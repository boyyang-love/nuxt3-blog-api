package helper

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"mime/multipart"
)

type CompressedImage struct {
	Buf    *bytes.Buffer
	Width  int
	Height int
	Hash   string
	Size   int64
}

func ResizeImage(fileHeader *multipart.FileHeader) (compressionImage *CompressedImage, err error) {
	imgFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(img multipart.File) {
		_ = img.Close()
	}(imgFile)

	img, imgType, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	cimg := resize.Resize(uint(bounds.Dx()), uint(bounds.Dy()), img, resize.Lanczos3)
	buffer := new(bytes.Buffer)
	hash, err := MakeImageFileHash(cimg, imgType)
	if err != nil {
		return nil, err
	}

	switch imgType {
	case "png":
		if err = png.Encode(buffer, cimg); err != nil {
			return nil, err
		}
	case "jpeg", "jpg":
		if err = jpeg.Encode(buffer, cimg, nil); err != nil {
			return nil, err
		}
	}

	reader := bytes.NewReader(buffer.Bytes())

	return &CompressedImage{
		Hash:   hash,
		Buf:    buffer,
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
		Size:   int64(reader.Size()),
	}, nil
}
