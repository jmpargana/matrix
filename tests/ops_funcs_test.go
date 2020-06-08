package matrix_test

import (
	"matrix"
	"testing"
)

func TestUnequal(t *testing.T) {
	a, b := matrix.New(5, 4), matrix.New(2, 3)
	if matrix.Equal(a, b) {
		t.Errorf("%v should not be equal %v", a, b)
	}
}

func TestFailMult(t *testing.T) {
	a, b := matrix.New(2, 3), matrix.New(2, 3)
	_, err := matrix.Mult(a, b)
	if err == nil {
		t.Errorf("not supposed to be able to multiply %v with %v", a, b)
	}
}

func TestAdd(t *testing.T) {
	a := matrix.New(6, 5)
	a.AddScalar(4)

	b := matrix.New(6, 5)
	b.AddScalar(4)

	c := matrix.New(6, 5)
	c.AddScalar(8)

	res, _ := matrix.Add(a, b)

	if !matrix.Equal(res, c) {
		t.Errorf("\n%s should be equal to \n%s", res, c)
	}
}

func TestFailAdd(t *testing.T) {
	a, b := matrix.New(1, 5), matrix.New(5, 1)
	_, err := matrix.Add(a, b)
	if err == nil {
		t.Errorf("not supposed to be able to add \n%s\nand\n%s", a, b)
	}
}

func TestSub(t *testing.T) {
	a := matrix.New(6, 5)
	a.AddScalar(4)

	b := matrix.New(6, 5)
	b.AddScalar(2)

	c := matrix.New(6, 5)
	c.AddScalar(2)

	res, _ := matrix.Sub(a, b)

	if !matrix.Equal(res, c) {
		t.Errorf("\n%s should be equal to \n%s", res, c)
	}
}

func TestFailSub(t *testing.T) {
	a, b := matrix.New(1, 5), matrix.New(5, 1)
	_, err := matrix.Sub(a, b)
	if err == nil {
		t.Errorf("not supposed to be able to add \n%s\nand\n%s", a, b)
	}
}

func TestMultVecsToMat(t *testing.T) {
	a, b := matrix.New(4, 1), matrix.New(1, 4)
	a.AddScalar(1)
	b.AddScalar(4)

	c := matrix.New(4, 4)
	c.AddScalar(4)

	res, _ := matrix.Mult(a, b)
	if !matrix.Equal(res, c) {
		t.Errorf("multipliying %v with %v should result in %v, instead got %v", a, b, c, res)
	}
}

func TestMultVecsTo1x1(t *testing.T) {
	a, b := matrix.New(1, 4), matrix.New(4, 1)
	a.AddScalar(1)
	b.AddScalar(4)

	c := matrix.New(1, 1)
	c.AddScalar(16)

	res, _ := matrix.Mult(a, b)
	if !matrix.Equal(res, c) {
		t.Errorf("multipliying %v with %v should result in %v, instead got %v", a, b, c, res)
	}
}

func TestMultSquare(t *testing.T) {
	a, b := matrix.New(4, 4), matrix.New(4, 4)
	a.AddScalar(1)
	b.AddScalar(4)

	c := matrix.New(4, 4)
	c.AddScalar(16)

	res, _ := matrix.Mult(a, b)
	if !matrix.Equal(res, c) {
		t.Errorf("multipliying %v with %v should result in %v, instead got %v", a, b, c, res)
	}
}

func TestTrans(t *testing.T) {
	for _, mats := range transposed {
		rhs := matrix.NewFrom(mats.rhs)
		lhs := matrix.NewFrom(mats.lhs)
		res, _ := matrix.Trans(rhs)

		if !matrix.Equal(res, lhs) {
			t.Errorf("%v should be the save as %v when transposed", rhs, lhs)
		}
	}
}
