package matrix_test

import (
	"testing"

	. "github.com/jmpargana/matrix"
)

func TestUnequal(t *testing.T) {
	a, b := New(5, 4), New(2, 3)
	if Equal(a, b) {
		t.Errorf("%v should not be equal %v", a, b)
	}
}

func TestFailMult(t *testing.T) {
	a, b := New(2, 3), New(2, 3)
	_, err := Mult(a, b)
	if err == nil {
		t.Errorf("not supposed to be able to multiply %v with %v", a, b)
	}
}

func TestAdd(t *testing.T) {
	a := New(6, 5)
	a.AddScalar(4)

	b := New(6, 5)
	b.AddScalar(4)

	c := New(6, 5)
	c.AddScalar(8)

	res, _ := Add(a, b)

	if !Equal(res, c) {
		t.Errorf("\n%s should be equal to \n%s", res, c)
	}
}

func TestFailAdd(t *testing.T) {
	a, b := New(1, 5), New(5, 1)
	_, err := Add(a, b)
	if err == nil {
		t.Errorf("not supposed to be able to add \n%s\nand\n%s", a, b)
	}
}

func TestSub(t *testing.T) {
	a := New(6, 5)
	a.AddScalar(4)

	b := New(6, 5)
	b.AddScalar(2)

	c := New(6, 5)
	c.AddScalar(2)

	res, _ := Sub(a, b)

	if !Equal(res, c) {
		t.Errorf("\n%s should be equal to \n%s", res, c)
	}
}

func TestFailSub(t *testing.T) {
	a, b := New(1, 5), New(5, 1)
	_, err := Sub(a, b)
	if err == nil {
		t.Errorf("not supposed to be able to add \n%s\nand\n%s", a, b)
	}
}

func TestMultVecsToMat(t *testing.T) {
	a, b := New(4, 1), New(1, 4)
	a.AddScalar(1)
	b.AddScalar(4)

	c := New(4, 4)
	c.AddScalar(4)

	res, _ := Mult(a, b)
	if !Equal(res, c) {
		t.Errorf("multipliying %v with %v should result in %v, instead got %v", a, b, c, res)
	}
}

func TestMultVecsTo1x1(t *testing.T) {
	a, b := New(1, 4), New(4, 1)
	a.AddScalar(1)
	b.AddScalar(4)

	c := New(1, 1)
	c.AddScalar(16)

	res, _ := Mult(a, b)
	if !Equal(res, c) {
		t.Errorf("multipliying %v with %v should result in %v, instead got %v", a, b, c, res)
	}
}

func TestMultSquare(t *testing.T) {
	a, b := New(4, 4), New(4, 4)
	a.AddScalar(1)
	b.AddScalar(4)

	c := New(4, 4)
	c.AddScalar(16)

	res, _ := Mult(a, b)
	if !Equal(res, c) {
		t.Errorf("multipliying %v with %v should result in %v, instead got %v", a, b, c, res)
	}
}

func TestTrans(t *testing.T) {
	for _, mats := range transposed {
		rhs := NewFrom(mats.rhs)
		lhs := NewFrom(mats.lhs)
		res, _ := Trans(rhs)

		if !Equal(res, lhs) {
			t.Errorf("%v should be the save as %v when transposed", rhs, lhs)
		}
	}
}

func TestMatMultVec(t *testing.T) {
	a, b, c := New(5, 1), New(3, 5), New(3, 1)
	a.AddScalar(1)
	b.AddScalar(1)
	c.AddScalar(5)

	got, _ := Mult(b, a)
	if !got.Equal(c) {
		t.Errorf("expected:\n%vgot:\n%v", c, got)
	}
}

func TestMatMultVecTrans(t *testing.T) {
	a, b, c := New(1, 5), New(5, 3), New(1, 3)
	a.AddScalar(1)
	b.AddScalar(1)
	c.AddScalar(5)

	got, _ := Mult(a, b)
	if !got.Equal(c) {
		t.Errorf("expected:\n%vgot:\n%v", c, got)
	}
}

func TestMatMultVec2(t *testing.T) {
	a := NewFrom([][]float64{
		{2, 4, 3, 4},
		{3, 2, 2, 3},
		{3, 1, -2, 3},
	})
	b := NewFrom([][]float64{
		{2},
		{-3},
		{-1},
		{3},
	})
	c := NewFrom([][]float64{
		{1},
		{7},
		{14},
	})

	got, err := Mult(a, b)
	if err != nil {
		t.Errorf("shouldn't fail here: %v", err)
	}

	if !got.Equal(c) {
		t.Errorf("\ngot:\n%vexpected:\n%v", got, c)
	}
}

func TestHadamardProd(t *testing.T) {
	for _, mats := range hadamardProdTest {
		a, b, c := NewFrom(mats.a), NewFrom(mats.b), NewFrom(mats.c)

		got1, err := HadamardProd(a, b)
		if err != nil {
			t.Errorf("wasn't supposed to fail here: %v", err)
		}

		got2, err := HadamardProd(b, a)
		if err != nil {
			t.Errorf("wasn't supposed to fail here: %v", err)
		}

		if !Equal(c, got1) {
			t.Errorf("got:\n%v\nexpected:\n%v\n", got1, c)
		}

		if !Equal(c, got2) {
			t.Errorf("got:\n%v\nexpected:\n%v\n", got2, c)
		}
	}
}

func TestHadamardProdFail(t *testing.T) {
	for _, mats := range hadamardProdTestInvalid {
		a, b := NewFrom(mats.a), NewFrom(mats.b)

		_, err := HadamardProd(a, b)
		if err == nil {
			t.Errorf("wasn supposed to fail here: %v", err)
		}
	}
}
