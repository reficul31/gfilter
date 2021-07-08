package gfilter

import (
	"image"
	"image/color/palette"
	"image/png"
	"os"

	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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
	ImageTypePaletted
)

// Pixel struct to contain RGBA components for images
type Pixel struct {
	r, g, b, a uint32
}

// ImageHandler interface for polymorphism
type ImageHandler interface {
	At(row, column int) (Pixel, error)
	Set(row, column int, px Pixel) error
	Mode() ImageType
	GetDimensions() image.Rectangle
	SaveImage(filePath string) error
}

// NRGBAImageHandler struct to handle NRGBA Image
type NRGBAImageHandler struct {
	Image     *image.NRGBA
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *NRGBAImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row < rect.Min.X && row > rect.Max.X) || (column < rect.Min.Y && column > rect.Max.Y) {
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
	return handler.ImageType
}

// GetDimensions returns the dimensions as a Rectangle for the image
func (handler *NRGBAImageHandler) GetDimensions() image.Rectangle {
	return handler.Image.Bounds()
}

// SaveImage saves the image to the filepath specified
func (handler *NRGBAImageHandler) SaveImage(filePath string) error {
	// Encode the grayscale image to the new file
	newfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newfile.Close()
	png.Encode(newfile, handler.Image)
	return nil
}

// NRGBA64ImageHandler struct to handle NRGBA64 Image
type NRGBA64ImageHandler struct {
	Image     *image.NRGBA64
	ImageType ImageType
}

// At retuns the Pixel at row and column of the image
func (handler *NRGBA64ImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row < rect.Min.X && row > rect.Max.X) || (column < rect.Min.Y && column > rect.Max.Y) {
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
	return handler.ImageType
}

// GetDimensions returns the dimensions as a Rectangle for the image
func (handler *NRGBA64ImageHandler) GetDimensions() image.Rectangle {
	return handler.Image.Bounds()
}

// SaveImage saves the image to the filepath specified
func (handler *NRGBA64ImageHandler) SaveImage(filePath string) error {
	// Encode the grayscale image to the new file
	newfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newfile.Close()
	png.Encode(newfile, handler.Image)
	return nil
}

// RGBAImageHandler struct to handle RGBA Image
type RGBAImageHandler struct {
	Image     *image.RGBA
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *RGBAImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row < rect.Min.X && row > rect.Max.X) || (column < rect.Min.Y && column > rect.Max.Y) {
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
	return handler.ImageType
}

// GetDimensions returns the dimensions as a Rectangle for the image
func (handler *RGBAImageHandler) GetDimensions() image.Rectangle {
	return handler.Image.Bounds()
}

// SaveImage saves the image to the filepath specified
func (handler *RGBAImageHandler) SaveImage(filePath string) error {
	// Encode the grayscale image to the new file
	newfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newfile.Close()
	png.Encode(newfile, handler.Image)
	return nil
}

// RGBA64ImageHandler struct to handle the RGBA64 Image
type RGBA64ImageHandler struct {
	Image     *image.RGBA64
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *RGBA64ImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row < rect.Min.X && row > rect.Max.X) || (column < rect.Min.Y && column > rect.Max.Y) {
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
	return handler.ImageType
}

// GetDimensions returns the dimensions as a Rectangle for the image
func (handler *RGBA64ImageHandler) GetDimensions() image.Rectangle {
	return handler.Image.Bounds()
}

// SaveImage saves the image to the filepath specified
func (handler *RGBA64ImageHandler) SaveImage(filePath string) error {
	// Encode the grayscale image to the new file
	newfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newfile.Close()
	png.Encode(newfile, handler.Image)
	return nil
}

// GrayImageHandler struct to handle the Gray Image
type GrayImageHandler struct {
	Image     *image.Gray
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *GrayImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row < rect.Min.X && row > rect.Max.X) || (column < rect.Min.Y && column > rect.Max.Y) {
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
	return handler.ImageType
}

// GetDimensions returns the dimensions as a Rectangle for the image
func (handler *GrayImageHandler) GetDimensions() image.Rectangle {
	return handler.Image.Bounds()
}

// SaveImage saves the image to the filepath specified
func (handler *GrayImageHandler) SaveImage(filePath string) error {
	// Encode the grayscale image to the new file
	newfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newfile.Close()
	png.Encode(newfile, handler.Image)
	return nil
}

// Gray16ImageHandler struct to handle the Gray16 Image
type Gray16ImageHandler struct {
	Image     *image.Gray16
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *Gray16ImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row < rect.Min.X && row > rect.Max.X) || (column < rect.Min.Y && column > rect.Max.Y) {
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
	return handler.ImageType
}

// GetDimensions returns the dimensions as a Rectangle for the image
func (handler *Gray16ImageHandler) GetDimensions() image.Rectangle {
	return handler.Image.Bounds()
}

// SaveImage saves the image to the filepath specified
func (handler *Gray16ImageHandler) SaveImage(filePath string) error {
	// Encode the grayscale image to the new file
	newfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newfile.Close()
	png.Encode(newfile, handler.Image)
	return nil
}

// Gray16ImageHandler struct to handle the Gray16 Image
type PalettedImageHandler struct {
	Image     *image.Paletted
	ImageType ImageType
}

// At returns the Pixel at row and column of the image
func (handler *PalettedImageHandler) At(row, column int) (Pixel, error) {
	rect := handler.Image.Bounds()
	if (row < rect.Min.X && row > rect.Max.X) || (column < rect.Min.Y && column > rect.Max.Y) {
		return Pixel{}, ErrRowColumnOutOfBounds
	}
	return ConvertColor(handler.Image.At(row, column)), nil
}

// Set sets the Pixel at row and column of the image
func (handler *PalettedImageHandler) Set(row, column int, px Pixel) error {
	setColor, err := GetColorFromPixel(px, handler.ImageType)
	if err != nil {
		return err
	}

	handler.Image.Set(row, column, setColor)
	return nil
}

// Mode returns the ImageType of the images
func (handler *PalettedImageHandler) Mode() ImageType {
	return handler.ImageType
}

// GetDimensions returns the dimensions as a Rectangle for the image
func (handler *PalettedImageHandler) GetDimensions() image.Rectangle {
	return handler.Image.Bounds()
}

// SaveImage saves the image to the filepath specified
func (handler *PalettedImageHandler) SaveImage(filePath string) error {
	newfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newfile.Close()
	png.Encode(newfile, handler.Image)
	return nil
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

	switch img.(type) {
	case *image.NRGBA:
		return &NRGBAImageHandler{
			Image:     img.(*image.NRGBA),
			ImageType: ImageTypeNRGBA,
		}, nil

	case *image.NRGBA64:
		return &NRGBA64ImageHandler{
			Image:     img.(*image.NRGBA64),
			ImageType: ImageTypeNRGBA64,
		}, nil

	case *image.RGBA:
		return &RGBAImageHandler{
			Image:     img.(*image.RGBA),
			ImageType: ImageTypeRGBA,
		}, nil

	case *image.RGBA64:
		return &RGBA64ImageHandler{
			Image:     img.(*image.RGBA64),
			ImageType: ImageTypeRGBA64,
		}, nil

	case *image.Gray:
		return &GrayImageHandler{
			Image:     img.(*image.Gray),
			ImageType: ImageTypeGray,
		}, nil

	case *image.Gray16:
		return &Gray16ImageHandler{
			Image:     img.(*image.Gray16),
			ImageType: ImageTypeGray16,
		}, nil

	case *image.Paletted:
		return &PalettedImageHandler{
			Image:     img.(*image.Paletted),
			ImageType: ImageTypePaletted,
		}, nil

	default:
		return nil, ErrUnsupportedImageFormat
	}
}

// New returns an ImageHandler with a new Image
// This handler will be used for the result
func New(imgT ImageType, dimensions image.Rectangle) (ImageHandler, error) {
	switch imgT {
	case ImageTypeNRGBA:
		return &NRGBAImageHandler{
			Image:     image.NewNRGBA(dimensions),
			ImageType: ImageTypeNRGBA,
		}, nil

	case ImageTypeNRGBA64:
		return &NRGBA64ImageHandler{
			Image:     image.NewNRGBA64(dimensions),
			ImageType: ImageTypeNRGBA64,
		}, nil

	case ImageTypeRGBA:
		return &RGBAImageHandler{
			Image:     image.NewRGBA(dimensions),
			ImageType: ImageTypeRGBA,
		}, nil

	case ImageTypeRGBA64:
		return &RGBA64ImageHandler{
			Image:     image.NewRGBA64(dimensions),
			ImageType: ImageTypeRGBA64,
		}, nil

	case ImageTypeGray:
		return &GrayImageHandler{
			Image:     image.NewGray(dimensions),
			ImageType: ImageTypeGray,
		}, nil

	case ImageTypeGray16:
		return &Gray16ImageHandler{
			Image:     image.NewGray16(dimensions),
			ImageType: ImageTypeGray16,
		}, nil

	case ImageTypePaletted:
		return &PalettedImageHandler{
			Image:     image.NewPaletted(dimensions, palette.Plan9),
			ImageType: ImageTypePaletted,
		}, nil
	default:
		return nil, ErrUnsupportedImageFormat
	}
}
