package gfilter

import (
	"fmt"
	"image"
	"image/draw"
)

// Apply is the convolution function used to apply the provided filter
// on the given image. It makes a copy and returns the filtered image
// without modifying the src image.
func Apply(filter *Filter, src image.Image) (draw.Image, error) {

}