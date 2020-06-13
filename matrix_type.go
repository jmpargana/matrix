// Package matrix is a go package that implements a simple matrix type.
// It has both methods for the type as well as functions for the most
// important operations: Mult, Add, Sub, MultScalar, AddScalar and HadamardProd.
// It uses concurrency only in the matrix multiplication, since its the only
// place where the overhead might actually be worth it.
package matrix

// Matrix is a matrix of floats. The size of the matrix is publicly available,
// while its data only via getters.
type Matrix struct {
	NumRows int
	NumCols int
	data    []float64
}
