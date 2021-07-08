package gfilter

// Kernel struct handles data for the filter kernel
type Kernel struct {
	values  []float32
	Sum     float32
	Rows    int
	Columns int
}

// CreateKernel creates a kernel object from the params
func CreateKernel(values []float32, rows, columns int) (*Kernel, error) {
	if rows == 0 || columns == 0 {
		return nil, ErrRowColumnCannotBeZero
	}

	var totalsum, possum float32 = 0, 0
	for i := 0; i < len(values); i = i + 1 {
		totalsum = totalsum + values[i]

		if values[i] > 0 {
			possum = possum + values[i]
		}
	}

	var sum float32
	if sum <= 0 {
		sum = possum
	} else {
		sum = totalsum
	}

	return &Kernel{
		values:  values,
		Sum:     sum,
		Rows:    rows,
		Columns: columns,
	}, nil
}

// At gets the value of the kernel at the specified row and column
func (k *Kernel) At(row, column int) (float32, error) {
	if row >= k.Rows || column >= k.Columns {
		return 0.0, ErrRowColumnOutOfBounds
	}

	return k.values[row*k.Columns+column], nil
}

// CalculateShift returns the amount the original image must be shifted
func (k *Kernel) CalculateShift() (int, int) {
	return int(k.Rows / 2), int(k.Columns / 2)
}
