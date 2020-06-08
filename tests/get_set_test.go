package matrix

import (
	"testing"
)

func TestRow(t *testing.T) {
	for _, mat := range rectMatrices {
		m := NewFrom(mat)

		for i := 0; i < m.NumRows; i++ {
			row, err := m.Row(i)
			if err != nil {
				t.Errorf("getting row with wrong index: %v", err)
			}

			if !compareSlices(row, mat[i]) {
				t.Errorf("row getter did not work. got: %v, expected: %v", row, mat[i])
			}
		}
	}
}

func TestCol(t *testing.T) {
	for _, mat := range rectMatrices {
		m := NewFrom(mat)

		for i := 0; i < m.NumCols; i++ {
			tmpCol := make([]float64, 0, len(mat))

			for _, row := range mat {
				tmpCol = append(tmpCol, row[i])
			}

			col, err := m.Col(i)
			if err != nil {
				t.Errorf("getting row with wrong index: %v", err)
			}

			if !compareSlices(col, tmpCol) {
				t.Errorf("row getter did not work. got: %v, expected: %v", col, tmpCol)
			}
		}
	}
}

func TestInvalidIndexing(t *testing.T) {
	for _, mat := range rectMatrices {
		m := NewFrom(mat)

		low := -1
		highRow := m.NumRows
		highCol := m.NumCols

		if _, err := m.Row(low); err == nil {
			t.Errorf("could get row with invalid index %d: %v", low, err)
		}

		if _, err := m.Row(highRow); err == nil {
			t.Errorf("could get row with invalid index %d: %v", highRow, err)
		}

		if _, err := m.Col(low); err == nil {
			t.Errorf("could get row with invalid index %d: %v", low, err)
		}

		if _, err := m.Col(highCol); err == nil {
			t.Errorf("could get row with invalid index %d: %v", highCol, err)
		}
	}
}

func TestGet(t *testing.T) {
	for _, mat := range rectMatrices {
		m := NewFrom(mat)

		for row := range mat {
			for col := range mat[row] {
				got, err := m.Get(row, col)
				if err != nil {
					t.Errorf("got error getting value at %d:%d", row, col)
				}

				if got != mat[row][col] {
					t.Errorf("wrong value at %d:%d", row, col)
				}
			}
		}
	}
}

func TestSet(t *testing.T) {
	for _, mat := range rectMatrices {
		m := NewFrom(mat)

		for row := range mat {
			for col := range mat[row] {
				var val float64 = 10 + float64(row+col)

				err := m.Set(row, col, val)

				if err != nil {
					t.Errorf("error setting at %d:%d: %v", row, col, err)
				}

				got, err := m.Get(row, col)
				if err != nil {
					t.Errorf("error getting at %d:%d: %v", row, col, err)
				}

				if got != val {
					t.Errorf("value wasn't saved in index %d:%d after settings to %f", row, col, val)
				}
			}
		}
	}
}

func TestInvalidGet(t *testing.T) {
	for _, mat := range rectMatrices {
		m := NewFrom(mat)

		low := -1
		highRow := m.NumRows
		highCol := m.NumCols

		if _, err := m.Get(low, low); err == nil {
			t.Errorf("should fail at Get in %d:%d in %dx%d matrix", low, low, m.NumRows, m.NumCols)
		}

		if _, err := m.Get(low, highCol); err == nil {
			t.Errorf("should fail at Get in %d:%d in %dx%d matrix", low, highCol, m.NumRows, m.NumCols)
		}

		if _, err := m.Get(highRow, low); err == nil {
			t.Errorf("should fail at Get in %d:%d in %dx%d matrix", highRow, low, m.NumRows, m.NumCols)
		}

		if _, err := m.Get(highRow, highCol); err == nil {
			t.Errorf("should fail at Get in %d:%d in %dx%d matrix", highRow, highCol, m.NumRows, m.NumCols)
		}
	}
}

func TestInvalidSet(t *testing.T) {
	for _, mat := range rectMatrices {
		m := NewFrom(mat)

		low := -1
		highRow := m.NumRows
		highCol := m.NumCols

		if err := m.Set(low, low, 0.0); err == nil {
			t.Errorf("should fail at setting in %d:%d in %dx%d matrix", low, low, m.NumRows, m.NumCols)
		}

		if err := m.Set(low, highCol, 0.0); err == nil {
			t.Errorf("should fail at setting in %d:%d in %dx%d matrix", low, highCol, m.NumRows, m.NumCols)
		}

		if err := m.Set(highRow, low, 0.0); err == nil {
			t.Errorf("should fail at setting in %d:%d in %dx%d matrix", highRow, low, m.NumRows, m.NumCols)
		}

		if err := m.Set(highRow, highCol, 0.0); err == nil {
			t.Errorf("should fail at setting in %d:%d in %dx%d matrix", highRow, highCol, m.NumRows, m.NumCols)
		}
	}
}
