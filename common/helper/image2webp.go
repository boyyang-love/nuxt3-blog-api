package helper

import (
	"bytes"
	"github.com/nickalie/go-webpbin"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

func Image2Webp(fileHeader *multipart.FileHeader, quality uint) (compressedImage *CompressedImage, err error) {

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
	}

	buf := new(bytes.Buffer)

	if err = webpbin.NewCWebP().
		Quality(quality).
		InputImage(img).
		Output(buf).
		Run(); err != nil {
		return nil, err
	}

	originReader := bytes.NewReader(originBuffer.Bytes())
	reader := bytes.NewReader(buf.Bytes())

	return &CompressedImage{
		Hash:       MakeImageFileHashByBytes(originBuffer.Bytes()),
		OriginBuf:  originBuffer,
		Buf:        buf,
		Width:      img.Bounds().Dx(),
		Height:     img.Bounds().Dy(),
		OriginSize: originReader.Size(),
		Size:       reader.Size(),
		OriginType: imgType,
		Type:       ".webp",
	}, nil

}
