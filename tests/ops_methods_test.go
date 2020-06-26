package matrix_test

import (
	"testing"

	"github.com/jmpargana/matrix"
	. "github.com/jmpargana/matrix"
)

func TestAddScalarMethod(t *testing.T) {
	for _, test := range addScalarMatrices {
		mat := NewFrom(test.in)
		expected := NewFrom(test.out)
		mat.AddScalar(test.step)

		if !Equal(mat, expected) {
			t.Errorf("add function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestInvalidAddScalarMethod(t *testing.T) {
	for _, test := range invalidAddScalarMatrices {
		mat := NewFrom(test.in)
		expected := NewFrom(test.out)
		mat.AddScalar(test.step)

		if Equal(mat, expected) {
			t.Errorf("add function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestMultScalarMethod(t *testing.T) {
	for _, test := range multScalarMatrices {
		mat := NewFrom(test.in)
		expected := NewFrom(test.out)
		mat.MultScalar(test.step)

		if !Equal(mat, expected) {
			t.Errorf("mult function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestInvalidMultScalarMethod(t *testing.T) {
	for _, test := range invalidMultScalarMatrices {
		mat := NewFrom(test.in)
		expected := NewFrom(test.out)
		mat.MultScalar(test.step)

		if Equal(mat, expected) {
			t.Errorf("mult function didn't work. expected: %v, got: %v", expected, mat)
		}
	}
}

func TestEqualMethod(t *testing.T) {
	for _, mat := range equalMatrices {
		rhs, lhs := NewFrom(mat.rhs), NewFrom(mat.lhs)
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
		rhs, lhs := NewFrom(mat.rhs), NewFrom(mat.lhs)
		if rhs.Equal(lhs) {
			t.Errorf("%v should not be equal to %v", rhs, lhs)
		}
	}
}

func TestUnequalElems(t *testing.T) {
	a, b := New(4, 4), New(4, 4)
	b.Set(1, 1, 1.00)

	if a.Equal(b) {
		t.Errorf("%v should not be equal to %v", a, b)
	}
}

func TestAddScalarMethod2(t *testing.T) {
	a, b := New(3, 4), New(3, 4)
	a.AddScalar(1)
	a.AddScalar(1)
	b.AddScalar(2)
	if !a.Equal(b) {
		t.Errorf("\n%s should be equal to \n%s", a, b)
	}
}

func TestMultScalarMethod2(t *testing.T) {
	a, b := New(1, 4), New(1, 4)
	a.AddScalar(2)
	a.MultScalar(2)
	b.AddScalar(4)
	if !a.Equal(b) {
		t.Errorf("\n%s should be equal to \n%s", a, b)
	}
}

func TestFailAddMethod1(t *testing.T) {
	a, b := New(3, 3), New(3, 4)
	err := a.Add(b)
	if err == nil {
		t.Errorf("should not be able to add %v to %v", a, b)
	}
}

func TestFailSubMethod2(t *testing.T) {
	a, b := New(4, 4), New(3, 4)
	err := a.Sub(b)
	if err == nil {
		t.Errorf("should not be able to subtract %v from %v", a, b)
	}
}
func TestFailAddMethod(t *testing.T) {
	a, b, c := New(3, 4), New(3, 4), New(3, 4)
	a.AddScalar(1)
	b.AddScalar(1)
	a.Add(b)

	if a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestFailSubMethod(t *testing.T) {
	a, b, c := New(3, 4), New(3, 4), New(3, 4)
	a.AddScalar(1)
	b.AddScalar(-1)
	a.Sub(b)

	if a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestAddMethod(t *testing.T) {
	a, b, c := New(3, 4), New(3, 4), New(3, 4)
	a.AddScalar(1)
	b.AddScalar(-1)
	a.Add(b)

	if !a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestSubMethod(t *testing.T) {
	a, b, c := New(3, 4), New(3, 4), New(3, 4)
	a.AddScalar(1)
	b.AddScalar(1)
	a.Sub(b)

	if !a.Equal(c) {
		t.Errorf("%v should be equal to %v", a, c)
	}
}

func TestMultMethodFail(t *testing.T) {
	a, b := New(4, 3), New(4, 3)
	err := a.Mult(b)
	if err == nil {
		t.Errorf("should be able to multiply %v with %v", a, b)
	}
}

func TestMultMethodFail2(t *testing.T) {
	a, b := New(4, 2), New(4, 3)
	err := a.Mult(b)
	if err == nil {
		t.Errorf("should be able to multiply %v with %v", a, b)
	}
}

func TestIsSquare(t *testing.T) {
	a, b := New(4, 4), New(2, 3)
	if !a.IsSquare() {
		t.Errorf("%v should be square", a)
	}
	if b.IsSquare() {
		t.Errorf("%v should not be square", b)
	}
}

func TestMultMethod(t *testing.T) {
	a, b, c := New(4, 4), New(4, 4), New(4, 4)
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
		rhs := NewFrom(mats.rhs)
		lhs := NewFrom(mats.lhs)
		rhs.Trans()

		if !rhs.Equal(lhs) {
			t.Errorf("%v should be the save as %v when transposed", rhs, lhs)
		}
	}
}

func TestEncodingDecoding(t *testing.T) {
	for name, tc := range encodingTestCases {
		t.Run(name, func(t *testing.T) {
			m := matrix.NewFrom(tc.m)
			mat := &matrix.Matrix{}

			data, err := m.MarshalBinary()
			if err != nil {
				t.Errorf("failed encoding with: %v", err)
			}

			err = mat.UnmarshalBinary(data)
			if err != nil {
				t.Errorf("failed decoding with: %v", err)
			}

			if !m.Equal(*mat) {
				t.Errorf("got\n%swant\n%s", mat, m)
			}
		})
	}
}
