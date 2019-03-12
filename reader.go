package gfilter

import (
	"os"
	"image"
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
)

// ImageType to add support for different images
type ImageType int

// Enumeration of image types
const (
	ImageTypeNRGBA = ImageType(iota)
	ImageTypeNRGBA64
	ImageTypeRGBA
	ImageTypeRGBA64
	ImageTypeGray
	ImageTypeGray16
)

// Pixel struct to contain RGBA components for images
type Pixel struct {
	r, g, b, a float32
}

// ImageHandler interface for polymorphism
type ImageHandler interface {
	At(row, column int) (Pixel, error)
	Set(row, column int, px Pixel) error
	Mode() ImageType
}

// NRGBAImageHandler struct to handle NRGBA Image
type NRGBAImageHandler struct {
	Image *image.NRGBA
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *NRGBAImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row >= rect.Min.X && row <= rect.Max.X) || (column >= rect.Min.Y && column <= rect.Max.Y) {
		return Pixel{}, ErrRowColumnOutOfBounds
	}
	return ConvertColor(handler.Image.At(row, column)), nil
}

// Set sets the Pixel at row and column of the image
func (handler *NRGBAImageHandler) Set(row, column int, px Pixel) error {
	setColor, err := GetColorFromPixel(px, handler.ImageType)
	if err != nil {
		return err
	}

	handler.Image.Set(row, column, setColor)
	return nil
}

// Mode returns the ImageType for the image
func (handler *NRGBAImageHandler) Mode() ImageType {
	return handler.ImageType;
}

// NRGBA64ImageHandler struct to handle NRGBA64 Image 
type NRGBA64ImageHandler struct {
	Image *image.NRGBA64
	ImageType ImageType
}

// At retuns the Pixel at row and column of the image
func (handler *NRGBA64ImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row >= rect.Min.X && row <= rect.Max.X) || (column >= rect.Min.Y && column <= rect.Max.Y) {
		return Pixel{}, ErrRowColumnOutOfBounds
	}
	return ConvertColor(handler.Image.At(row, column)), nil
}

// Set sets the Pixel at row and column of the image
func (handler *NRGBA64ImageHandler) Set(row, column int, px Pixel) error {
	setColor, err := GetColorFromPixel(px, handler.ImageType)
	if err != nil {
		return err
	}

	handler.Image.Set(row, column, setColor)
	return nil
}

// Mode returns the ImageType of the images
func (handler *NRGBA64ImageHandler) Mode() ImageType {
	return handler.ImageType;
}

// RGBAImageHandler struct to handle RGBA Image
type RGBAImageHandler struct {
	Image *image.RGBA
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *RGBAImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row >= rect.Min.X && row <= rect.Max.X) || (column >= rect.Min.Y && column <= rect.Max.Y) {
		return Pixel{}, ErrRowColumnOutOfBounds
	}
	return ConvertColor(handler.Image.At(row, column)), nil
}

// Set sets the Pixel at row and column of the image
func (handler *RGBAImageHandler) Set(row, column int, px Pixel) error {
	setColor, err := GetColorFromPixel(px, handler.ImageType)
	if err != nil {
		return err
	}

	handler.Image.Set(row, column, setColor)
	return nil
}

// Mode returns the ImageType of the images
func (handler *RGBAImageHandler) Mode() ImageType {
	return handler.ImageType;
}

// RGBA64ImageHandler struct to handle the RGBA64 Image
type RGBA64ImageHandler struct {
	Image *image.RGBA64
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *RGBA64ImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row >= rect.Min.X && row <= rect.Max.X) || (column >= rect.Min.Y && column <= rect.Max.Y) {
		return Pixel{}, ErrRowColumnOutOfBounds
	}
	return ConvertColor(handler.Image.At(row, column)), nil
}

// Set sets the Pixel at row and column of the image
func (handler *RGBA64ImageHandler) Set(row, column int, px Pixel) error {
	setColor, err := GetColorFromPixel(px, handler.ImageType)
	if err != nil {
		return err
	}

	handler.Image.Set(row, column, setColor)
	return nil
}

// Mode returns the ImageType of the images
func (handler *RGBA64ImageHandler) Mode() ImageType {
	return handler.ImageType;
}

// GrayImageHandler struct to handle the Gray Image
type GrayImageHandler struct {
	Image *image.Gray
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *GrayImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row >= rect.Min.X && row <= rect.Max.X) || (column >= rect.Min.Y && column <= rect.Max.Y) {
		return Pixel{}, ErrRowColumnOutOfBounds
	}
	return ConvertColor(handler.Image.At(row, column)), nil
}

// Set sets the Pixel at row and column of the image
func (handler *GrayImageHandler) Set(row, column int, px Pixel) error {
	setColor, err := GetColorFromPixel(px, handler.ImageType)
	if err != nil {
		return err
	}

	handler.Image.Set(row, column, setColor)
	return nil
}

// Mode returns the ImageType of the images
func (handler *GrayImageHandler) Mode() ImageType {
	return handler.ImageType;
}

// Gray16ImageHandler struct to handle the Gray16 Image
type Gray16ImageHandler struct {
	Image *image.Gray16
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *Gray16ImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row >= rect.Min.X && row <= rect.Max.X) || (column >= rect.Min.Y && column <= rect.Max.Y) {
		return Pixel{}, ErrRowColumnOutOfBounds
	}
	return ConvertColor(handler.Image.At(row, column)), nil
}

// Set sets the Pixel at row and column of the image
func (handler *Gray16ImageHandler) Set(row, column int, px Pixel) error {
	setColor, err := GetColorFromPixel(px, handler.ImageType)
	if err != nil {
		return err
	}

	handler.Image.Set(row, column, setColor)
	return nil
}

// Mode returns the ImageType of the images
func (handler *Gray16ImageHandler) Mode() ImageType {
	return handler.ImageType;
}

// ReadImage reads the image from the specified path
// it will return the interface ImageHandler of the 
// correct type specified in the image
func ReadImage(path string) (ImageHandler, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	switch imgT := img.(type) {
	case *image.NRGBA:
		return &NRGBAImageHandler{
			Image: img.(*image.NRGBA),
			ImageType: ImageTypeNRGBA,
		}, nil

	case *image.NRGBA64:
		return &NRGBA64ImageHandler{
			Image: img.(*image.NRGBA64),
			ImageType: ImageTypeNRGBA64,
		}, nil

	case *image.RGBA:
		return &RGBAImageHandler{
			Image: img.(*image.RGBA),
			ImageType: ImageTypeRGBA,
		}, nil

	case *image.RGBA64:
		return &RGBA64ImageHandler{
			Image: img.(*image.RGBA64),
			ImageType: ImageTypeRGBA64,
		}, nil

	case *image.Gray:
		return &GrayImageHandler{
			Image: img.(*image.Gray),
			ImageType: ImageTypeGray,
		}, nil

	case *image.Gray16:
		return &Gray16ImageHandler{
			Image: img.(*image.Gray16),
			ImageType: ImageTypeGray16,
		}, nil

	default:
		return nil, ErrUnsupportedImageFormat
	}
}