package gfilter

import "math"

// Apply is the convolution function used to apply the provided filter
// on the given image. It makes a copy and returns the filtered image
// without modifying the src image.
func Apply(filter *Kernel, src ImageHandler) (ImageHandler, error) {
	result, err := New(ImageTypeNRGBA64, src.GetDimensions())
	if err != nil {
		return nil, err
	}

	xShift, yShift := filter.CalculateShift()
	for i := src.GetDimensions().Min.X; i < src.GetDimensions().Max.X-int(filter.Rows); i = i + 1 {
		for j := src.GetDimensions().Min.Y; j < src.GetDimensions().Max.Y-int(filter.Columns); j = j + 1 {
			var rSum, gSum, bSum float64 = 0, 0, 0
			for row := 0; row < int(filter.Rows); row = row + 1 {
				for column := 0; column < int(filter.Columns); column = column + 1 {
					pixel, err := src.At(i+row, j+column)
					if err != nil {
						return nil, err
					}
					multiplier, err := filter.At(row, column)
					if err != nil {
						return nil, err
					}
					rSum = rSum + float64(pixel.r)*float64(multiplier)
					gSum = gSum + float64(pixel.g)*float64(multiplier)
					bSum = bSum + float64(pixel.b)*float64(multiplier)
				}
			}

			oldPixel, err := src.At(i+xShift, j+yShift)
			if err != nil {
				return nil, err
			}
			newPixel := Pixel{r: uint32(math.Max(rSum, 0) / float64(filter.Sum)), g: uint32(math.Max(gSum, 0) / float64(filter.Sum)), b: uint32(math.Max(bSum, 0) / float64(filter.Sum)), a: uint32(oldPixel.a)}
			result.Set(i+xShift, j+yShift, newPixel)
		}
	}

	return result, nil
}
