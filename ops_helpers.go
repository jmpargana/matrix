package matrix

import (
	"errors"
	"fmt"
)

func assertMatchingSizes(lhsRow, lhsCol, rhsRow, rhsCol int) error {
	if lhsRow != rhsRow || lhsCol != rhsCol {
		return errors.New(fmt.Sprintf("can't add non matching matrices: mat: %dx%d, other: %dx%d", lhsRow, lhsCol, rhsRow, rhsCol))
	}
	return nil
}

func assertMatchingMult(lhsCol, rhsRow int) error {
	if lhsCol != rhsRow {
		return errors.New(fmt.Sprintf("can't multiply non matching matrices: mat: %d columns and other: %d rows", lhsCol, rhsRow))
	}
	return nil
}

func dot(row, col []float64) (r float64) {
	for i := range row {
		r += row[i] * col[i]
	}
	return
}
