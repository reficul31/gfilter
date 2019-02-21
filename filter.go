package gfilter

// Kernel struct handles data for the filter kernel
type Kernel struct {
	values  []uint32
	rows    uint32
	columns uint32
}

// CreateKernel creates a kernel object from the params
func CreateKernel(values []uint32, rows, columns uint32) (*Kernel, error) {
	if rows == 0 || columns == 0 {
		return nil, ErrRowColumnCannotBeZero
	}
	return &Kernel{
		values:  values,
		rows:    rows,
		columns: columns,
	}, nil
}
