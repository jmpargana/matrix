package matrix

import (
	"bytes"
	"encoding/gob"
	"errors"
	"sync"
)

// AddScalar adds a value to all elements in matrix.
func (m *Matrix) AddScalar(val float64) {
	for i := range m.data {
		m.data[i] += val
	}
}

// Add adds a value to first matrix and saves result in its data.
// Both matrices should have the same dimensions.
func (m *Matrix) Add(other Matrix) error {
	if err := assertMatchingSizes(m.NumRows, m.NumCols, other.NumRows, other.NumCols); err != nil {
		return err
	}

	for i := range m.data {
		m.data[i] += other.data[i]
	}

	return nil
}

// Add subtracts a value to first matrix and saves result in its data.
// Both matrices should have the same dimensions.
func (m *Matrix) Sub(other Matrix) error {
	if err := assertMatchingSizes(m.NumRows, m.NumCols, other.NumRows, other.NumCols); err != nil {
		return err
	}

	for i := range m.data {
		m.data[i] -= other.data[i]
	}

	return nil
}

// MultScalar multiplies all elements of matrix by value.
func (m *Matrix) MultScalar(val float64) {
	for i := range m.data {
		m.data[i] *= val
	}
}

// Mult multiplies two matrices and saves result in first one.
// The two matrices must look the same, otherwise the shape of the first needs
// to change.
func (m *Matrix) Mult(other Matrix) error {
	if err := assertMatchingSizes(m.NumRows, m.NumCols, other.NumRows, other.NumCols); err != nil {
		return err
	}

	if !m.IsSquare() {
		return errors.New("the multiply method should only work on square matrices")
	}

	wg := sync.WaitGroup{}
	tmp := New(m.NumRows, m.NumCols)

	for row := 0; row < m.NumRows; row++ {
		for col := 0; col < other.NumCols; col++ {
			wg.Add(1)

			lhsRow, _ := m.Row(row)
			rhsCol, _ := other.Col(col)

			go func(i, j int) {
				tmp.Set(i, j, dot(lhsRow, rhsCol))
				wg.Done()
			}(row, col)
		}
	}
	wg.Wait()
	m.data = tmp.data

	return nil
}

// Trans transposes the matrix in place.
func (m *Matrix) Trans() error {
	tmp := make([]float64, 0, m.NumRows*m.NumCols)

	for i := 0; i < m.NumCols; i++ {

		col, err := m.Col(i)
		if err != nil {
			return err
		}

		tmp = append(tmp, col...)
	}
	m.data = tmp
	m.NumRows, m.NumCols = m.NumCols, m.NumRows

	return nil
}

// Equal compares the instance matrix with another.
func (m *Matrix) Equal(other Matrix) bool {
	if m.NumRows != other.NumRows || m.NumCols != other.NumCols {
		return false
	}

	for i := range m.data {
		if m.data[i] != other.data[i] {
			return false
		}
	}
	return true
}

// IsSquare compares the number of columns with the number of rows.
func (m *Matrix) IsSquare() bool {
	return m.NumRows == m.NumCols
}

// MarshalBinary is the method needed to implement the BinaryMarshaler interface.
// With it one can call gob.Encode on the Matrix struct.
func (m *Matrix) MarshalBinary() ([]byte, error) {
	w := wrapMatrix{m.NumRows, m.NumCols, m.data}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalBinary is the method needed to implement the BinaryUnmarshaler interface.
// With it one can call gob.Decode on the Matrix struct.
func (m *Matrix) UnmarshalBinary(data []byte) error {
	w := wrapMatrix{}

	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(&w); err != nil {
		return err
	}

	m.NumRows = w.NumRows
	m.NumCols = w.NumCols
	m.data = w.Data

	return nil
}
