package matrix

import (
	"errors"
	"fmt"
)

// Row returns a slice of the row on the nth index of the matrix.
func (m *Matrix) Row(row int) ([]float64, error) {
	if row < 0 || row >= m.NumRows {
		return nil, errors.New(fmt.Sprintf("tried to index %d row from %dx%d matrix", row, m.NumRows, m.NumCols))
	}
	return m.data[row*m.NumCols : (row+1)*m.NumCols], nil
}

// Col returns a slice of the column on the nth index of the matrix.
func (m *Matrix) Col(col int) ([]float64, error) {
	if col < 0 || col >= m.NumCols {
		return nil, errors.New(fmt.Sprintf("tried to index %d row col %dx%d matrix", col, m.NumRows, m.NumCols))
	}

	result := make([]float64, 0, m.NumRows)

	for i := 0; i < m.NumRows; i++ {
		result = append(result, m.data[i*m.NumCols+col])
	}

	return result, nil
}

func (m *Matrix) Get(row, col int) (float64, error) {
	if row < 0 || col < 0 || row >= m.NumRows || col >= m.NumCols {
		return 0, errors.New(fmt.Sprintf("tried to set at wrong index %d:%d in %dx%d matrix", row, col, m.NumRows, m.NumCols))
	}

	return m.data[row*m.NumCols+col], nil
}

func (m *Matrix) Set(row, col int, val float64) error {
	if row < 0 || col < 0 || row >= m.NumRows || col >= m.NumCols {
		return errors.New(fmt.Sprintf("tried to set at wrong index %d:%d in %dx%d matrix", row, col, m.NumRows, m.NumCols))
	}

	m.data[row*m.NumCols+col] = val

	return nil
}
