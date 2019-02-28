package gfilter

import "errors"

var (
	// ErrRowColumnCannotBeZero thrown when either row or columns or both is zero
	ErrRowColumnCannotBeZero = errors.New("Row and Column cannot be zero")
	// ErrRowColumnOutOfBounds is thrown when row and/or column is out of bounds for the defined kernel
	ErrRowColumnOutOfBounds = errors.New("Row and/or Column is out of bounds")
	// ErrUnsupportedImageFormat is thrown when the image is of unsupported format
	ErrUnsupportedImageFormat = errors.New("The image format is unsupported")
)
