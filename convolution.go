package gfilter

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
			var rSum, gSum, bSum uint32 = 0, 0, 0
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
					rSum = rSum + pixel.r*uint32(multiplier)
					gSum = gSum + pixel.g*uint32(multiplier)
					bSum = bSum + pixel.b*uint32(multiplier)
				}
			}

			oldPixel, err := src.At(i+xShift, j+yShift)
			if err != nil {
				return nil, err
			}
			newPixel := Pixel{r: rSum / uint32(filter.Sum), g: gSum / uint32(filter.Sum), b: bSum / uint32(filter.Sum), a: uint32(oldPixel.a)}
			result.Set(i+xShift, j+yShift, newPixel)
		}
	}

	return result, nil
}
