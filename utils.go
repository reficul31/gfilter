package gfilter

import "color"

// RoundD helps round down from float64 to uint
func RoundD(val float32) uint {
	if val > 0 {
		return uint(val+1.0)
	}
	return uint(val-1.0)
}

// ConvertColor converts color.Color to Pixel
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

// GetColorFromPixel converts Pixel to color.Color
func GetColorFromPixel(px Pixel, imageType ImageType) (color.Color, error) {
	switch(imageType) {
	case ImageTypeNRGBA:
		return color.NRGBA{R: RoundD(px.r), G: RoundD(px.g), B: RoundD(px.b), A: RoundD(px.a)}, nil

	case ImageTypeNRGBA64:
		return color.NRGBA64{R: RoundD(px.r), G: RoundD(px.g), B: RoundD(px.b), A: RoundD(px.a)}, nil

	case ImageTypeRGBA:
		return color.RGBA{R: RoundD(px.r), G: RoundD(px.g), B: RoundD(px.b), A: RoundD(px.a)}, nil

	case ImageTypeRGBA64:
		return color.RGBA64{R: RoundD(px.r), G: RoundD(px.g), B: RoundD(px.b), A: RoundD(px.a)}, nil

	case ImageTypeGray:
		return color.Gray{uint8(0.299*px.r + 0.587*px.g + 0.114*px.b / 256)}, nil

	case ImageTypeGray16:
		return color.Gray16{uint8(0.299*px.r + 0.587*px.g + 0.114*px.b / 256)}, nil

	default:
		return nil, ErrUnsupportedImageFormat
	}
}