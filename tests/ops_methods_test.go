package matrix_test

import (
	"matrix"
	"testing"
)

func TestAddScalarMethod(t *testing.T) {
	for _, test := range addScalarMatrices {
		mat := matrix.NewFrom(test.in)
		expected := matrix.NewFrom(test.out)
		mat.AddScalar(test.step)

		if !matrix.Equal(mat, expected) {
			t.Errorf("add function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestInvalidAddScalarMethod(t *testing.T) {
	for _, test := range invalidAddScalarMatrices {
		mat := matrix.NewFrom(test.in)
		expected := matrix.NewFrom(test.out)
		mat.AddScalar(test.step)

		if matrix.Equal(mat, expected) {
			t.Errorf("add function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestMultScalarMethod(t *testing.T) {
	for _, test := range multScalarMatrices {
		mat := matrix.NewFrom(test.in)
		expected := matrix.NewFrom(test.out)
		mat.MultScalar(test.step)

		if !matrix.Equal(mat, expected) {
			t.Errorf("mult function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestInvalidMultScalarMethod(t *testing.T) {
	for _, test := range invalidMultScalarMatrices {
		mat := matrix.NewFrom(test.in)
		expected := matrix.NewFrom(test.out)
		mat.MultScalar(test.step)

		if matrix.Equal(mat, expected) {
			t.Errorf("mult function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestEqualMethod(t *testing.T) {
	for _, mat := range equalMatrices {
		rhs, lhs := matrix.NewFrom(mat.rhs), matrix.NewFrom(mat.lhs)
		if !rhs.Equal(lhs) {
			t.Errorf("%v should be equal to %v", rhs, lhs)
		}
	}
}

func TestEqualMethod2(t *testing.T) {
	for _, mat := range moreEqual {
		if !mat.a.Equal(mat.b) {
			t.Errorf("%v should be equal to %v", mat.a, mat.b)
		}
		if !mat.b.Equal(mat.c) {
			t.Errorf("%v should be equal to %v", mat.b, mat.c)
		}
		if !mat.a.Equal(mat.c) {
			t.Errorf("%v should be equal to %v", mat.a, mat.c)
		}
		if !mat.b.Equal(mat.a) {
			t.Errorf("%v should be equal to %v", mat.b, mat.a)
		}
		if !mat.c.Equal(mat.b) {
			t.Errorf("%v should be equal to %v", mat.c, mat.b)
		}
		if !mat.c.Equal(mat.a) {
			t.Errorf("%v should be equal to %v", mat.c, mat.a)
		}
		if !mat.a.Equal(mat.a) {
			t.Errorf("%v should be equal to %v", mat.a, mat.a)
		}
		if !mat.b.Equal(mat.b) {
			t.Errorf("%v should be equal to %v", mat.b, mat.b)
		}
		if !mat.c.Equal(mat.c) {
			t.Errorf("%v should be equal to %v", mat.c, mat.c)
		}
	}
}

func TestUnequalMethod(t *testing.T) {
	for _, mat := range unequalMatrices {
		rhs, lhs := matrix.NewFrom(mat.rhs), matrix.NewFrom(mat.lhs)
		if rhs.Equal(lhs) {
			t.Errorf("%v should not be equal to %v", rhs, lhs)
		}
	}
}

func TestUnequalElems(t *testing.T) {
	a, b := matrix.New(4, 4), matrix.New(4, 4)
	b.Set(1, 1, 1.00)

	if a.Equal(b) {
		t.Errorf("%v should not be equal to %v", a, b)
	}
}

func TestAddScalarMethod2(t *testing.T) {
	a, b := matrix.New(3, 4), matrix.New(3, 4)
	a.AddScalar(1)
	a.AddScalar(1)
	b.AddScalar(2)
	if !a.Equal(b) {
		t.Errorf("\n%s should be equal to \n%s", a, b)
	}
}

func TestMultScalarMethod2(t *testing.T) {
	a, b := matrix.New(1, 4), matrix.New(1, 4)
	a.AddScalar(2)
	a.MultScalar(2)
	b.AddScalar(4)
	if !a.Equal(b) {
		t.Errorf("\n%s should be equal to \n%s", a, b)
	}
}

func TestFailAddMethod1(t *testing.T) {
	a, b := matrix.New(3, 3), matrix.New(3, 4)
	err := a.Add(b)
	if err == nil {
		t.Errorf("should not be able to add %v to %v", a, b)
	}
}

func TestFailSubMethod2(t *testing.T) {
	a, b := matrix.New(4, 4), matrix.New(3, 4)
	err := a.Sub(b)
	if err == nil {
		t.Errorf("should not be able to subtract %v from %v", a, b)
	}
}
func TestFailAddMethod(t *testing.T) {
	a, b, c := matrix.New(3, 4), matrix.New(3, 4), matrix.New(3, 4)
	a.AddScalar(1)
	b.AddScalar(1)
	a.Add(b)

	if a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestFailSubMethod(t *testing.T) {
	a, b, c := matrix.New(3, 4), matrix.New(3, 4), matrix.New(3, 4)
	a.AddScalar(1)
	b.AddScalar(-1)
	a.Sub(b)

	if a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestAddMethod(t *testing.T) {
	a, b, c := matrix.New(3, 4), matrix.New(3, 4), matrix.New(3, 4)
	a.AddScalar(1)
	b.AddScalar(-1)
	a.Add(b)

	if !a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestSubMethod(t *testing.T) {
	a, b, c := matrix.New(3, 4), matrix.New(3, 4), matrix.New(3, 4)
	a.AddScalar(1)
	b.AddScalar(1)
	a.Sub(b)

	if !a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestMultMethodFail(t *testing.T) {
	a, b := matrix.New(4, 3), matrix.New(4, 3)
	err := a.Mult(b)
	if err == nil {
		t.Errorf("should be able to multiply %v with %v", a, b)
	}
}

func TestMultMethodFail2(t *testing.T) {
	a, b := matrix.New(4, 2), matrix.New(4, 3)
	err := a.Mult(b)
	if err == nil {
		t.Errorf("should be able to multiply %v with %v", a, b)
	}
}

func TestIsSquare(t *testing.T) {
	a, b := matrix.New(4, 4), matrix.New(2, 3)
	if !a.IsSquare() {
		t.Errorf("%v should be square", a)
	}
	if b.IsSquare() {
		t.Errorf("%v should not be square", b)
	}
}

func TestMultMethod(t *testing.T) {
	a, b, c := matrix.New(4, 4), matrix.New(4, 4), matrix.New(4, 4)
	a.AddScalar(1)
	b.AddScalar(1)
	c.AddScalar(4)
	a.Mult(b)

	if !a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestTransMethod(t *testing.T) {
	for _, mats := range transposed {
		rhs := matrix.NewFrom(mats.rhs)
		lhs := matrix.NewFrom(mats.lhs)
		rhs.Trans()

		if !rhs.Equal(lhs) {
			t.Errorf("%v should be the save as %v when transposed", rhs, lhs)
		}
	}
}
