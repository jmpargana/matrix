package matrix

import (
	"fmt"
	"math/rand"
	"time"
)

// New creates a new empty matrix of a given size.
func New(rows, cols int) Matrix {
	assertValidSize(rows, cols)
	return createWithData(rows, cols, make([]float64, rows*cols))
}

// NewSquare creates an empty square matrix of a given size.
func NewSquare(size int) Matrix {
	assertValidSize(size, size)
	return createWithData(size, size, make([]float64, size*size))
}

// NewRandom creates a matrix of a given size and fills the vectors with random
// floating point number ranging from 0..1
func NewRandom(rows, cols int) Matrix {
	assertValidSize(rows, cols)

	data := make([]float64, 0, rows*cols)
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range data {
		data[i] = gen.Float64()
	}

	return createWithData(rows, cols, data)
}

// NewFrom creates a matrix instance from a slice of slices
func NewFrom(data [][]float64) Matrix {
	rows, cols := len(data), len(data[0])

	assertValidSize(rows, cols)
	assertAllRowsSameSize(data, cols)

	return createWithData(rows, cols, sliceJoin(data))
}

// NewFromVector creates a matrix from long vector. It just makes sure it can have a
// rectangular shape
func NewFromVec(rows, cols int, data []float64) Matrix {
	assertValidSize(rows, cols)
	if rows*cols != len(data) {
		panic(fmt.Sprintf("tried to create %dx%d matrix with %v", rows, cols, data))
	}

	return createWithData(rows, cols, data)
}

// createWithData is called from all constructors and returns a struct
func createWithData(rows, cols int, data []float64) Matrix {
	return Matrix{
		NumRows: rows,
		NumCols: cols,
		data:    data,
	}
}

func assertValidSize(rows, cols int) {
	if rows < 1 || cols < 1 {
		panic(fmt.Sprintf("tried to create matrix with %d rows and %d columns", rows, cols))
	}
}

func assertAllRowsSameSize(data [][]float64, cols int) {
	for _, r := range data {
		if len(r) != cols {
			panic(fmt.Sprintf("tried to create matrix from slice of different sized slices: %v", data))
		}
	}
}

func sliceJoin(data [][]float64) []float64 {
	result := make([]float64, 0, len(data)*len(data[0]))

	for _, r := range data {
		result = append(result, r...)
	}

	return result
}
