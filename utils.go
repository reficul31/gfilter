package gfilter

import "image/color"

// ConvertColor converts color.Color to Pixel
func ConvertColor(c color.Color) Pixel {
	r, g, b, a := c.RGBA()
	if a == 0 {
		return Pixel{0, 0, 0, 0}
	}

	return Pixel{r, g, b, a}
}

// GetColorFromPixel converts Pixel to color.Color
func GetColorFromPixel(px Pixel, imageType ImageType) (color.Color, error) {
	switch imageType {
	case ImageTypePaletted:
		return color.RGBA{R: uint8(px.r), G: uint8(px.g), B: uint8(px.b), A: uint8(px.a)}, nil

	case ImageTypeNRGBA:
		return color.NRGBA{R: uint8(px.r), G: uint8(px.g), B: uint8(px.b), A: uint8(px.a)}, nil

	case ImageTypeNRGBA64:
		return color.NRGBA64{R: uint16(px.r), G: uint16(px.g), B: uint16(px.b), A: uint16(px.a)}, nil

	case ImageTypeRGBA:
		return color.RGBA{R: uint8(px.r), G: uint8(px.g), B: uint8(px.b), A: uint8(px.a)}, nil

	case ImageTypeRGBA64:
		return color.RGBA64{R: uint16(px.r), G: uint16(px.g), B: uint16(px.b), A: uint16(px.a)}, nil

	case ImageTypeGray:
		return color.Gray{uint8((px.r + px.g + px.b) / 3)}, nil

	case ImageTypeGray16:
		return color.Gray16{uint16((px.r + px.g + px.b) / 3)}, nil

	default:
		return nil, ErrUnsupportedImageFormat
	}
}
