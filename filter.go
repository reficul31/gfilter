package gfilter

// Kernel struct handles data for the filter kernel
type Kernel struct {
	values  []float32
	Rows    uint32
	Columns uint32
}

// CreateKernel creates a kernel object from the params
func CreateKernel(values []float32, rows, columns uint32) (*Kernel, error) {
	if rows == 0 || columns == 0 {
		return nil, ErrRowColumnCannotBeZero
	}
	return &Kernel{
		values:  values,
		Rows:    rows,
		Columns: columns,
	}, nil
}

// At gets the value of the kernel at the specified row and column
func (k *Kernel)At(row, column uint32) (float32, error) {
	if row >= k.Rows || column >= k.Columns {
		return 0.0, ErrRowColumnOutOfBounds
	}

	return k.values[row*k.Columns + column], nil
}

// CalculateShift returns the amount the original image must be shifted
func (k *Kernel)CalculateShift() (int, int) {
	return int(k.Rows/2), int(k.Columns/2)
}