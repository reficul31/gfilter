package gfilter

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"
	"color"
	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
)

type ImageType int

// 
const (
	ImageTypeNRGBA = ImageType(iota)
	ImageTypeNRGBA64
	ImageTypeRGBA
	ImageTypeRGBA64
	ImageTypeGray
	ImageTypeGray16
)

type Pixel struct {
	r, g, b, a float32
}

type ImageHandler interface {
	Image image.Image
	ImageType ImageType
	At(row, column int) (Pixel, error)
	Set(row, column int, px Pixel)
	Mode() int
}

type NRGBAImageHandler struct {
	Image image.Image
	ImageType ImageType
}

func (handler *NRGBAImageHandler) At(row, column int) (Pixel, error) {
	return ConvertColor(handler.Image.At(row, Column))
}

func (handler *NRGBAImageHandler) Set(row, column int, px Pixel) {
	color.Color setColor = GetColorFromPixel(px, handler.ImageType)
	handler.Image.set(row, column, setColor)
}

func (handler *NRGBAImageHandler) Mode() int {
	return handler.ImageType;
}

type NRGBA64ImageHandler struct {
	Image image.Image
	ImageType ImageType
}

func (handler *NRGBA64ImageHandler) At(row, column int) (Pixel, error) {
	return ConvertColor(handler.Image.At(row, column))
}

func (handler *NRGBA64ImageHandler) Set(row, column int, px Pixel) {
	color.Color setColor = GetColorFromPixel(px, handler.ImageType)
	handler.Image.set(row, column, setColor)
}

func (handler *NRGBA64ImageHandler) Mode() int {
	return handler.ImageType;
}

type RGBAImageHandler struct {
	Image image.Image
	ImageType ImageType
}

func (handler *RGBAImageHandler) At(row, column int) (Pixel, error) {
	return ConvertColor(handler.Image.At(row, Column))
}

func (handler *RGBAImageHandler) Set(row, column int, px Pixel) {
	color.Color setColor = GetColorFromPixel(px, handler.ImageType)
	handler.Image.set(row, column, setColor)
}

func (handler *RGBAImageHandler) Mode() int {
	return handler.ImageType;
}

type RGBA64ImageHandler struct {
	Image image.Image
	ImageType ImageType
}

func (handler *RGBA64ImageHandler) At(row, column int) (Pixel, error) {
	return ConvertColor(handler.Image.At(row, column))
}

func (handler *RGBA64ImageHandler) Set(row, column int, px Pixel) {
	color.Color setColor = GetColorFromPixel(px, handler.ImageType); 
	handler.Image.set(row, column, setColor);
}

func (handler *RGBA64ImageHandler) Mode() int {
	return handler.ImageType;
}

type Gray16ImageHandler struct {
	Image image.ImageType
	ImageType ImageType
}

func (handler *Gray16ImageHandler) At(row, column int) (Pixel, error) {
	return ConvertColor(handler.Image.At(row, column))
}

func (handler *Gray16ImageHandler) Set(row, column int, px Pixel) {
	color.Color setColor = GetColorFromPixel(px, handler.ImageType); 
	handler.Image.set(row, column, setColor);
}

func (handler *Gray16ImageHandler) Mode() int {
	return handler.ImageType;
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

	switch imgT := img.(type); imgT {
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

	default:
		return nil, ErrUnsupportedImageFormat
	}
}