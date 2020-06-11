package matrix_test

import . "github.com/jmpargana/matrix"

type scalarStruct struct {
	in   [][]float64
	step float64
	out  [][]float64
}

type compareMatrices struct {
	rhs [][]float64
	lhs [][]float64
}

type matrixOps struct {
	a, b, c Matrix
}

type matrixOpsFloats struct {
	a, b, c [][]float64
}

var matrixSize = [...][]int{
	{8, 4},
	{3, 3},
	{5, 2},
}

var invalidMatrix = [...][]int{
	{0, 4},
	{4, 0},
	{0, 0},
	{-1, 3},
	{4, -2},
	{-5, -2},
}

var vectorToMatrixOdd = [...][]float64{
	{3, 2, 1},
	{3, 4, 2, 2, 3},
}

var invalidSlices = [...][][]float64{
	{
		{2, 3},
		{1},
	},
	{
		{3},
		{1, 3, 2},
	},
}

var validSlices = [...][][]float64{
	{
		{2, 3},
		{1, 3},
	},
	{
		{3, 2, 3},
		{1, 3, 2},
	},
}

var vectorToMatrixEven = [...][]float64{
	{2, 3, 1, 4},
	{2, 3, 4, 1, 2, 3},
	{2, 3, 4, 1, 2, 3, 4, 5},
	{2, 3, 4, 2, 3, 2, 3, 3, 2, 3},
}

var rectMatrices = [...][][]float64{
	{
		{1, 2, 3, 4},
		{8, 7, 6, 5},
		{9, 10, 11, 12},
	},
	{
		{2, 3, 4, 2, 3, 4},
		{3, 2, 3, 4, 2, 3},
		{2, 3, 4, 5, 2, 3},
	},
	{
		{2, 3, 4, 2, 3, 4},
		{3, 2, 3, 4, 2, 3},
		{2, 3, 4, 5, 2, 3},
		{1, 1, 1, 1, 1, 11},
		{1, 3, 4, 5, 2, 1},
	},
	{
		{2, 3, 4, 2},
		{3, 2, 3, 4},
	},
	{
		{2, 3},
		{3, 2},
		{2, 3},
	},
}

var addScalarMatrices = []scalarStruct{
	{
		in: [][]float64{
			{1, 2},
			{3, 4},
		},
		step: 1,
		out: [][]float64{
			{2, 3},
			{4, 5},
		},
	},
	{
		in: [][]float64{
			{1, 2},
			{3, 4},
			{6, 6},
		},
		step: 3,
		out: [][]float64{
			{4, 5},
			{6, 7},
			{9, 9},
		},
	},
	{
		in: [][]float64{
			{1, 2, 3},
			{3, 4, 5},
		},
		step: -1,
		out: [][]float64{
			{0, 1, 2},
			{2, 3, 4},
		},
	},
}

var invalidAddScalarMatrices = []scalarStruct{
	{
		in: [][]float64{
			{1, 2},
			{3, 4},
		},
		step: 2,
		out: [][]float64{
			{2, 3},
			{4, 5},
		},
	},
	{
		in: [][]float64{
			{1, 2},
			{3, 4},
			{6, 6},
		},
		step: 1,
		out: [][]float64{
			{4, 5},
			{6, 7},
			{9, 9},
		},
	},
	{
		in: [][]float64{
			{1, 2, 3},
			{3, 4, 5},
		},
		step: -1,
		out: [][]float64{
			{1, 1, 2},
			{2, 3, 4},
		},
	},
}

var multScalarMatrices = []scalarStruct{
	{
		in: [][]float64{
			{1, 2},
			{3, 4},
		},
		step: 2,
		out: [][]float64{
			{2, 4},
			{6, 8},
		},
	},
	{
		in: [][]float64{
			{1, 2},
			{3, 4},
			{6, 6},
		},
		step: 0.5,
		out: [][]float64{
			{0.5, 1},
			{1.5, 2},
			{3, 3},
		},
	},
}

var invalidMultScalarMatrices = []scalarStruct{
	{
		in: [][]float64{
			{1, 2},
			{3, 4},
		},
		step: 2,
		out: [][]float64{
			{1, 4},
			{6, 8},
		},
	},
	{
		in: [][]float64{
			{1, 2},
			{1, 4},
			{6, 6},
		},
		step: 1 / 2,
		out: [][]float64{
			{1 / 2, 1},
			{3 / 2, 2},
			{1, 3},
		},
	},
}

var equalMatrices = []compareMatrices{
	{
		rhs: [][]float64{
			{3, 4},
			{3, 4},
		},
		lhs: [][]float64{
			{3, 4},
			{3, 4},
		},
	},
	{
		rhs: [][]float64{
			{3, 4},
		},
		lhs: [][]float64{
			{3, 4},
		},
	},
	{
		rhs: [][]float64{
			{3},
			{3},
		},
		lhs: [][]float64{
			{3},
			{3},
		},
	},
	{
		rhs: [][]float64{
			{3},
		},
		lhs: [][]float64{
			{3},
		},
	},
	{
		rhs: [][]float64{
			{3, 4, 3},
			{3, 4, 3},
		},
		lhs: [][]float64{
			{3, 4, 3},
			{3, 4, 3},
		},
	},
}

var unequalMatrices = []compareMatrices{
	{
		rhs: [][]float64{
			{3, 4},
		},
		lhs: [][]float64{
			{3, 4},
			{3, 4},
		},
	},
	{
		rhs: [][]float64{
			{3},
		},
		lhs: [][]float64{
			{3, 4},
		},
	},
	{
		rhs: [][]float64{
			{3, 4},
			{3, 6},
		},
		lhs: [][]float64{
			{3},
			{3},
		},
	},
	{
		rhs: [][]float64{
			{3},
		},
		lhs: [][]float64{
			{3},
			{4},
		},
	},
	{
		rhs: [][]float64{
			{3, 3},
			{3, 3},
		},
		lhs: [][]float64{
			{3, 4, 3},
			{3, 4, 3},
		},
	},
}

var moreEqual = []matrixOps{
	{
		New(4, 3),
		New(4, 3),
		New(4, 3),
	},
	{
		New(2, 3),
		New(2, 3),
		New(2, 3),
	},
	{
		New(3, 3),
		New(3, 3),
		New(3, 3),
	},
	{
		New(4, 4),
		New(4, 4),
		New(4, 4),
	},
}

var transposed = []compareMatrices{
	{
		[][]float64{
			{1, 2},
			{3, 4},
			{5, 5},
		},
		[][]float64{
			{1, 3, 5},
			{2, 4, 5},
		},
	},
	{
		[][]float64{
			{1, 2},
			{3, 4},
		},
		[][]float64{
			{1, 3},
			{2, 4},
		},
	},
	{
		[][]float64{
			{1, 2, 3},
		},
		[][]float64{
			{1},
			{2},
			{3},
		},
	},
}

var hadamardProdTest = []matrixOpsFloats{
	{
		a: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
			{-2, -4, -2, -3},
		},
		b: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
			{-2, -4, -2, -3},
		},
		c: [][]float64{
			{1, 4, 9, 16},
			{9, 1, 16, 4},
			{4, 16, 4, 9},
		},
	},
	{
		a: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
			{-2, -4, -2, -3},
		},
		b: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
			{2, 4, 2, 3},
		},
		c: [][]float64{
			{1, 4, 9, 16},
			{9, 1, 16, 4},
			{-4, -16, -4, -9},
		},
	},
	{
		a: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
			{-2, -4, -2, -3},
			{-2, -4, -2, -3},
		},
		b: [][]float64{
			{1, 2, 3, 4},
			{0, 1, 4, 2},
			{0, -4, -2, -3},
			{0, 0, -2, -3},
		},
		c: [][]float64{
			{1, 4, 9, 16},
			{0, 1, 16, 4},
			{0, 16, 4, 9},
			{0, 0, 4, 9},
		},
	},
}

var hadamardProdTestInvalid = []matrixOpsFloats{
	{
		a: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
		},
		b: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
			{-2, -4, -2, -3},
		},
		c: [][]float64{
			{1, 4, 9, 16},
			{9, 1, 16, 4},
			{4, 16, 4, 9},
		},
	},
	{
		a: [][]float64{
			{1, 2, 3},
			{3, 1, 4},
			{-2, -4, -3},
		},
		b: [][]float64{
			{1, 2, 3, 4},
			{3, 1, 4, 2},
			{2, 4, 2, 3},
		},
		c: [][]float64{
			{1, 4, 9, 16},
			{9, 1, 16, 4},
			{-4, -16, -4, -9},
		},
	},
	{
		a: [][]float64{
			{-2, -4, -2, -3},
		},
		b: [][]float64{
			{1, 2, 3, 4},
			{0, 1, 4, 2},
		},
		c: [][]float64{
			{1, 4, 9, 16},
			{0, 1, 16, 4},
			{0, 16, 4, 9},
			{0, 0, 4, 9},
		},
	},
}
