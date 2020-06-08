package matrix

import "sync"

func Add(lhs, rhs Matrix) (Matrix, error) {
	if err := assertMatchingSizes(lhs.NumRows, lhs.NumCols, rhs.NumRows, lhs.NumCols); err != nil {
		return Matrix{}, err
	}

	data := make([]float64, rhs.NumRows*rhs.NumCols)

	for i := range lhs.data {
		data[i] = lhs.data[i] + rhs.data[i]
	}

	return NewFromVec(rhs.NumRows, rhs.NumCols, data), nil
}

func Sub(lhs, rhs Matrix) (Matrix, error) {
	if err := assertMatchingSizes(lhs.NumRows, lhs.NumCols, rhs.NumRows, lhs.NumCols); err != nil {
		return Matrix{}, err
	}

	data := make([]float64, rhs.NumRows*rhs.NumCols)

	for i := range lhs.data {
		data[i] = lhs.data[i] - rhs.data[i]
	}

	return NewFromVec(rhs.NumRows, rhs.NumCols, data), nil
}

func Mult(lhs, rhs Matrix) (Matrix, error) {
	if err := assertMatchingMult(lhs.NumCols, rhs.NumRows); err != nil {
		return Matrix{}, err
	}
	result := New(rhs.NumCols, lhs.NumRows)
	wg := sync.WaitGroup{}

	for row := 0; row < lhs.NumRows; row++ {
		for col := 0; col < rhs.NumCols; col++ {
			wg.Add(1)

			lhsRow, _ := lhs.Row(row)
			rhsCol, _ := rhs.Col(col)

			go func(i, j int) {
				result.Set(i, j, dot(lhsRow, rhsCol))
				wg.Done()
			}(row, col)
		}
	}
	wg.Wait()

	return result, nil
}

// Trans returns a new matrix which represents the transposed version of the first.
func Trans(m Matrix) (Matrix, error) {
	data := make([]float64, 0, m.NumRows*m.NumCols)

	for i := 0; i < m.NumCols; i++ {

		col, err := m.Col(i)
		if err != nil {
			return Matrix{}, err
		}

		data = append(data, col...)
	}

	return NewFromVec(m.NumCols, m.NumRows, data), nil
}

// Equal compares to matrices for equality.
func Equal(rhs, lhs Matrix) bool {
	if rhs.NumRows != lhs.NumRows || rhs.NumCols != lhs.NumCols {
		return false
	}

	for i := range rhs.data {
		if rhs.data[i] != lhs.data[i] {
			return false
		}
	}
	return true
}
