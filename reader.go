package gfilter

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
)

// 
const (
	ImageType int iota
	ImageTypeNRGBA
	ImageTypeNRGBA64
	ImageTypeRGBA
	ImageTypeRGBA64
	ImageTypeYCbCr
	ImageTypeGray
	ImageTypeGray16
	ImageTypePaletted
)

type Pixel struct {
	r, g, b, a float32
}

type ImageHandler interface {
	Image image.Image
	ImageType int
	At(row, column int) (Pixel, error)
	Set(row, column int, px Pixel) (error)
	Mode() (int, error)
}

type NRGBAImageHandler struct {
	Image image.Image
	ImageType int
}

func (handler *NRGBAImageHandler) At(row, column int) Pixel {
	return ConvertColor(handler.Image.At(row, Column))
}

func (handler *NRGBAImageHandler) Set(row, column int, px Pixel) {
	
}

func ConvertColor(c color.Color) Pixel {
	rColor, gColor, bColor, aColor := c.RGBA()
	if(aColor == 0) {
		return Pixel{0, 0, 0, 0}
	} else {
		r := float32(rColor) / float32(aColor)
		g := float32(gColor) / float32(aColor)
		b := float32(bColor) / float32(aColor)
		a := float32(aColor) / float32(aColor)
		return Pixel{r, g, b, a}
	}
}


func ReadImage(path string) (*ImageHandler, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	switch img := img.(type) {
	case *image.NRGBA:
		return &NRGBAImageHandler{
			Image: img,
			ImageType: ImageTypeNRGBA,
		}, nil

	case *image.NRGBA64:
		return &NRGBA64ImageHandler{
			Image: img,
			ImageType: ImageTypeNRGBA64,
		}

	case *image.RGBA:
		return &RGBAImageHandler{
			Image: img,
			ImageType: ImageTypeRGBA,
		}, nil

	case *image.RGBA64:
		return &RGBA64ImageHandler{
			Image: img,
			ImageType: ImageTypeRGBA64,
		}, nil

	case *image.Gray:
		return &GrayImageHandler{
			Image: img,
			ImageType: ImageTypeGray,
		}, nil

	case *image.Gray16:
		return &Gray16ImageHandler{
			Image: img,
			ImageType: ImageTypeGray16,
		}, nil

	case *image.YCbCr:
		return &YCbCrImageHandler{
			Image: img,
			ImageType: ImageTypeYCbCr,
		}, nil

	case *image.Paletted:
		return PalettedImageHandler{
			Image: img,
			ImageType: ImageTypePaletted,
		}, nil

	default:
		return nil, ErrUnsupportedImageFormat
	}
}