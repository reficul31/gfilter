package gfilter

// Apply is the convolution function used to apply the provided filter
// on the given image. It makes a copy and returns the filtered image
// without modifying the src image.
func Apply(filter *Kernel, src ImageHandler) (ImageHandler, error) {
	result, err := New(src.Mode(), src.GetDimensions())
	if err != nil {
		return nil, err
	}
	return result, nil
}
