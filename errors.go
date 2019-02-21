package gfilter

import "errors"

var (
	// ErrRowColumnCannotBeZero thrown when either row or columns or both is zero
	ErrRowColumnCannotBeZero = errors.New("Row and Column cannot be zero")
)
