package helper

import (
	"bytes"
	"github.com/chai2010/webp"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"mime/multipart"
)

func ImageToWebp(fileHeader *multipart.FileHeader, quality float32) (compressedImage *CompressedImage, err error) {

	imageFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(imageFile)

	img, imgType, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	originBuffer := new(bytes.Buffer)
	switch imgType {
	case "png":
		if err = png.Encode(originBuffer, img); err != nil {
			return nil, err
		}
	case "jpeg", "jpg":
		if err = jpeg.Encode(originBuffer, img, nil); err != nil {
			return nil, err
		}
	case "webp":
		if err = webp.Encode(originBuffer, img, nil); err != nil {
			return nil, err
		}
	}

	imgByte, err := webp.EncodeRGBA(img, quality)
	if err != nil {
		return nil, err
	}

	hash := MakeImageFileHashByBytes(imgByte)

	reader := bytes.NewReader(imgByte)
	oriReader := bytes.NewReader(originBuffer.Bytes())

	return &CompressedImage{
		OriginBuf:  originBuffer,
		Buf:        bytes.NewBuffer(imgByte),
		Width:      img.Bounds().Dx(),
		Height:     img.Bounds().Dy(),
		Size:       reader.Size(),
		OriginSize: oriReader.Size(),
		Hash:       hash,
		OriginType: imgType,
		Type:       ".webp",
	}, nil
}
