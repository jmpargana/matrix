package matrix

import (
	"fmt"
)

func (m Matrix) String() (s string) {
	for i := 0; i < m.NumRows; i++ {
		for j := 0; j < m.NumCols; j++ {
			elem, _ := m.Get(i, j)
			s += fmt.Sprintf("%.2f", elem)
			s += " "
		}
		s += "\n"
	}
	return
}
