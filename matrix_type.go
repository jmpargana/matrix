package matrix

// Matrix is a matrix of floats. The size of the matrix is publicly available,
// while its data only via getters.
type Matrix struct {
	NumRows int
	NumCols int
	data    []float64
}
