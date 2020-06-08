package matrix_test

import (
	. "github.com/jmpargana/matrix"
	"testing"
)

func TestValidSlices(t *testing.T) {
	for _, ss := range validSlices {
		m := NewFrom(ss)

		if m.NumRows != len(ss) || m.NumCols != len(ss[0]) {
			t.Errorf("invalid construction of %v", ss)
		}
	}
}

func TestInvalidSlices(t *testing.T) {
	for _, ss := range invalidSlices {
		assertPanic(t, func() { NewFrom(ss) })
	}
}

func TestVectorToMatrixEven(t *testing.T) {
	for _, v := range vectorToMatrixEven {
		size := len(v)

		var hori = NewFromVec(1, size, v)
		var vert = NewFromVec(size, 1, v)
		var rectH = NewFromVec(2, size/2, v)
		var rectV = NewFromVec(size/2, 2, v)
		_, _, _, _ = hori, vert, rectH, rectV
	}
}

func TestVectorToMatrix(t *testing.T) {
	for _, v := range vectorToMatrixOdd {
		size := len(v)

		var validVecHorizontal = NewFromVec(1, size, v)
		var validVecVertical = NewFromVec(size, 1, v)
		_, _ = validVecVertical, validVecHorizontal

		assertPanic(t, func() { NewFromVec(size/2, 2, v) })
		assertPanic(t, func() { NewFromVec(2, size/2, v) })
	}
}

func TestInvalidMatrixConstructors(t *testing.T) {
	for _, m := range invalidMatrix {
		rows, cols := m[0], m[1]
		assertPanic(t, func() { New(rows, cols) })

		if rows < 1 {
			assertPanic(t, func() { NewSquare(rows) })
		}
		if cols < 1 {
			assertPanic(t, func() { NewSquare(cols) })
		}
	}
}

func TestMatrixNew(t *testing.T) {
	for i := range matrixSize {
		rows, cols := matrixSize[i][0], matrixSize[i][1]

		var m = New(rows, cols)

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

		var m = NewSquare(rows)

		if r := m.NumRows; r != rows {
			t.Errorf("failed create struct with %d rows: %d", rows, r)
		}

		if c := m.NumCols; c != rows {
			t.Errorf("failed to create struct with %d cols: %d", rows, c)
		}
	}
}
