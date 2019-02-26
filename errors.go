package gfilter

import "errors"

var (
	// ErrRowColumnCannotBeZero thrown when either row or columns or both is zero
	ErrRowColumnCannotBeZero = errors.New("Row and Column cannot be zero")
	// ErrRowColumnOutOfBounds is out of bounds for the defined kernel
	ErrRowColumnOutOfBounds = errors.New("Row and/or Column is out of bounds")
)
