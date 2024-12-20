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
	OriginBuf  *bytes.Buffer
	Buf        *bytes.Buffer
	Width      int
	Height     int
	Hash       string
	OriginSize int64
	Size       int64
	OriginType string
	Type       string
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

	hash, err := MakeImageFileHash(cimg, imgType)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	originBuffer := new(bytes.Buffer)
	switch imgType {
	case "png":
		if err = png.Encode(buffer, cimg); err != nil {
			return nil, err
		}
		if err = png.Encode(originBuffer, img); err != nil {
			return nil, err
		}
	case "jpeg", "jpg":
		if err = jpeg.Encode(buffer, cimg, nil); err != nil {
			return nil, err
		}
		if err = jpeg.Encode(originBuffer, img, nil); err != nil {
			return nil, err
		}
	}

	reader := bytes.NewReader(buffer.Bytes())

	return &CompressedImage{
		Hash:      hash,
		OriginBuf: originBuffer,
		Buf:       buffer,
		Width:     bounds.Dx(),
		Height:    bounds.Dy(),
		Size:      reader.Size(),
		Type:      imgType,
	}, nil
}
