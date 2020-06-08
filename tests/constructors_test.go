package matrix_test

import (
	"matrix"
	"testing"
)

func TestValidSlices(t *testing.T) {
	for _, ss := range validSlices {
		m := matrix.NewFrom(ss)

		if m.NumRows != len(ss) || m.NumCols != len(ss[0]) {
			t.Errorf("invalid construction of %v", ss)
		}
	}
}

func TestInvalidSlices(t *testing.T) {
	for _, ss := range invalidSlices {
		assertPanic(t, func() { matrix.NewFrom(ss) })
	}
}

func TestVectorToMatrixEven(t *testing.T) {
	for _, v := range vectorToMatrixEven {
		size := len(v)

		var hori = matrix.NewFromVec(1, size, v)
		var vert = matrix.NewFromVec(size, 1, v)
		var rectH = matrix.NewFromVec(2, size/2, v)
		var rectV = matrix.NewFromVec(size/2, 2, v)
		_, _, _, _ = hori, vert, rectH, rectV
	}
}

func TestVectorToMatrix(t *testing.T) {
	for _, v := range vectorToMatrixOdd {
		size := len(v)

		var validVecHorizontal = matrix.NewFromVec(1, size, v)
		var validVecVertical = matrix.NewFromVec(size, 1, v)
		_, _ = validVecVertical, validVecHorizontal

		assertPanic(t, func() { matrix.NewFromVec(size/2, 2, v) })
		assertPanic(t, func() { matrix.NewFromVec(2, size/2, v) })
	}
}

func TestInvalidMatrixConstructors(t *testing.T) {
	for _, m := range invalidMatrix {
		rows, cols := m[0], m[1]
		assertPanic(t, func() { matrix.New(rows, cols) })

		if rows < 1 {
			assertPanic(t, func() { matrix.NewSquare(rows) })
		}
		if cols < 1 {
			assertPanic(t, func() { matrix.NewSquare(cols) })
		}
	}
}

func TestMatrixNew(t *testing.T) {
	for i := range matrixSize {
		rows, cols := matrixSize[i][0], matrixSize[i][1]

		var m = matrix.New(rows, cols)

		if r := m.NumRows; r != rows {
			t.Errorf("failed create struct with %d rows: %d", rows, r)
		}

		if c := m.NumCols; c != cols {
			t.Errorf("failed to create struct with %d cols: %d", cols, c)
		}
	}
}

func TestMatrixSquare(t *testing.T) {
	for i := range matrixSize {
		rows := matrixSize[i][0]

		var m = matrix.NewSquare(rows)

		if r := m.NumRows; r != rows {
			t.Errorf("failed create struct with %d rows: %d", rows, r)
		}

		if c := m.NumCols; c != rows {
			t.Errorf("failed to create struct with %d cols: %d", rows, c)
		}
	}
}
