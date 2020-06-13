package matrix

import (
	"fmt"
)

func (m Matrix) String() (s string) {
	s += fmt.Sprintf("%dx%d", m.NumRows, m.NumCols)
	s += "\n\t"

	for i := 0; i < m.NumCols; i++ {
		s += fmt.Sprintf("col %d\t", i+1)
	}
	s += "\n"

	for i := 0; i < m.NumRows; i++ {
		s += fmt.Sprintf("row %d: ", i+1)
		for j := 0; j < m.NumCols; j++ {
			elem, _ := m.Get(i, j)
			s += fmt.Sprintf("%.5f", elem)
			s += " "
		}
		s += "\n"
	}
	return
}
