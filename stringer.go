package matrix

import (
	"fmt"
)

func (m Matrix) String() (s string) {
	s += fmt.Sprintf("%dx%d", m.NumRows, m.NumCols)
	for i := 0; i < m.NumRows; i++ {
		for j := 0; j < m.NumCols; j++ {
			elem, _ := m.Get(i, j)
			s += fmt.Sprintf("%.5f", elem)
			s += " "
		}
		s += "\n"
	}
	return
}
