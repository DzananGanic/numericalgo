package numericalgo_test

import (
	"fmt"
	"testing"

	"github.com/DzananGanic/numericalgo"
	"github.com/stretchr/testify/assert"
)

func TestCreateMatrix(t *testing.T) {
	// Creating empty matrix
	_ = make(numericalgo.Matrix, 2, 2)

	// With predefined values
	_ = numericalgo.Matrix{
		{1, 2},
		{3, 4},
	}
}

func TestCompareMatrices(t *testing.T) {

	cases := []struct {
		matrix1 numericalgo.Matrix
		matrix2 numericalgo.Matrix
		isEqual bool
	}{
		// Basic
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			isEqual: true,
		},
		// Wrong Dimensions
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
			},
			matrix2: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			isEqual: false,
		},
		// Wrong values
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{3, 5},
			},
			matrix2: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			isEqual: false,
		},
		// Nil passed
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{3, 5},
			},
			matrix2: nil,
			isEqual: false,
		},
	}

	for _, c := range cases {
		isEqual := c.matrix1.IsEqual(c.matrix2)
		assert.Equal(t, c.isEqual, isEqual)
	}
}

func TestMatrixAddColumnAt(t *testing.T) {
	cases := []struct {
		matrix         numericalgo.Matrix
		column         numericalgo.Vector
		index          int
		expectedResult numericalgo.Matrix
		expectedError  error
	}{
		{
			matrix: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			column: numericalgo.Vector{1, 1},
			index:  0,
			expectedResult: numericalgo.Matrix{
				{1, 1, 2},
				{1, 3, 4},
			},
			expectedError: nil,
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			column: numericalgo.Vector{1, 1},
			index:  1,
			expectedResult: numericalgo.Matrix{
				{1, 1, 2},
				{3, 1, 4},
			},
			expectedError: nil,
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			column: numericalgo.Vector{1, 1},
			index:  2,
			expectedResult: numericalgo.Matrix{
				{1, 2, 1},
				{3, 4, 1},
			},
			expectedError: nil,
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			column:         numericalgo.Vector{1, 1, 4},
			index:          0,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Column dimensions must match"),
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			column:         numericalgo.Vector{1, 1},
			index:          -1,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be less than 0"),
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			column:         numericalgo.Vector{1, 1},
			index:          3,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be greater than number of columns + 1"),
		},
	}

	for _, c := range cases {
		result, err := c.matrix.AddColumnAt(c.index, c.column)
		assert.Equal(t, result, c.expectedResult)
		assert.Equal(t, err, c.expectedError)
	}
}

func TestAddMatrices(t *testing.T) {
	cases := []struct {
		matrix1       numericalgo.Matrix
		matrix2       numericalgo.Matrix
		result        numericalgo.Matrix
		expectedError error
	}{
		// Basic
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: numericalgo.Matrix{
				{4, 3},
				{2, 1},
			},
			result: numericalgo.Matrix{
				{5, 5},
				{5, 5},
			},
			expectedError: nil,
		},
		// Wrong dimensions
		// {
		// 	matrix1: numericalgo.Matrix{
		// 		{1, 2},
		// 		{3, 4},
		// 	},
		// 	matrix2: numericalgo.Matrix{
		// 		{4, 3},
		// 	},
		// 	result:        nil,
		// 	expectedError: fmt.Errorf("Matrix dimensions must match"),
		// },
		// Adding two nils
		{
			matrix1:       nil,
			matrix2:       nil,
			result:        nil,
			expectedError: fmt.Errorf("Matrices cannot be nil"),
		},
	}

	for _, c := range cases {
		additionResult, err := c.matrix1.Add(c.matrix2)
		isCorrect := c.result.IsEqual(additionResult)
		assert.Equal(t, isCorrect, true)
		assert.Equal(t, err, c.expectedError)
	}

}

func TestSubtractMatrices(t *testing.T) {
	cases := []struct {
		matrix1       numericalgo.Matrix
		matrix2       numericalgo.Matrix
		result        numericalgo.Matrix
		expectedError error
	}{
		// Basic
		{
			matrix1: numericalgo.Matrix{
				{10, 5},
				{3, 1},
			},
			matrix2: numericalgo.Matrix{
				{1, 1},
				{1, 1},
			},
			result: numericalgo.Matrix{
				{9, 4},
				{2, 0},
			},
			expectedError: nil,
		},
		// With negative result
		{
			matrix1: numericalgo.Matrix{
				{3, 2},
				{3, 1},
			},
			matrix2: numericalgo.Matrix{
				{4, 3},
				{4, 2},
			},
			result: numericalgo.Matrix{
				{-1, -1},
				{-1, -1},
			},
			expectedError: nil,
		},
		// Wrong dimensions
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: numericalgo.Matrix{
				{4, 3},
			},
			result:        nil,
			expectedError: fmt.Errorf("Matrix dimensions must match"),
		},
		// Adding two nils
		{
			matrix1:       nil,
			matrix2:       nil,
			result:        nil,
			expectedError: fmt.Errorf("Matrices cannot be nil"),
		},
	}

	for _, c := range cases {
		additionResult, err := c.matrix1.Subtract(c.matrix2)
		isCorrect := c.result.IsEqual(additionResult)
		assert.Equal(t, isCorrect, true)
		assert.Equal(t, err, c.expectedError)
	}

}

func TestGetColumnAt(t *testing.T) {
	cases := []struct {
		matrix         numericalgo.Matrix
		expectedResult numericalgo.Vector
		i              int
		expectedError  error
	}{
		{
			matrix: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              1,
			expectedResult: numericalgo.Vector{2, 5},
			expectedError:  nil,
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              -5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be negative"),
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be greater than the length"),
		},
	}

	for _, c := range cases {
		column, err := c.matrix.GetColumnAt(c.i)
		assert.Equal(t, column, c.expectedResult)
		assert.Equal(t, err, c.expectedError)
	}

}

func TestGetRowAt(t *testing.T) {
	cases := []struct {
		matrix         numericalgo.Matrix
		expectedResult numericalgo.Vector
		i              int
		expectedError  error
	}{
		{
			matrix: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              1,
			expectedResult: numericalgo.Vector{4, 5, 6},
			expectedError:  nil,
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              -5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be negative"),
		},
		{
			matrix: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			i:              5,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Index cannot be greater than the length"),
		},
	}

	for _, c := range cases {
		column, err := c.matrix.GetRowAt(c.i)
		assert.Equal(t, column, c.expectedResult)
		assert.Equal(t, err, c.expectedError)
	}

}

func TestTransposeMatrix(t *testing.T) {
	cases := []struct {
		matrix        numericalgo.Matrix
		transposed    numericalgo.Matrix
		expectedError error
	}{
		// Basic
		{
			matrix: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			transposed: numericalgo.Matrix{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			expectedError: nil,
		},
		// Turning it the other way around
		{
			matrix: numericalgo.Matrix{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			transposed: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedError: nil,
		},
		// With one dimension only
		{
			matrix: numericalgo.Matrix{
				{1, 4},
			},
			transposed: numericalgo.Matrix{
				{1},
				{4},
			},
			expectedError: nil,
		},
		// Inconsistent dimensions
		// {
		// 	matrix: numericalgo.Matrix{
		// 		{1, 4},
		// 		{2},
		// 	},
		// 	transposed:    nil,
		// 	expectedError: fmt.Errorf("Inconsistent dimensions"),
		// },
	}

	for _, c := range cases {
		transposed, err := c.matrix.Transpose()
		isCorrect := c.transposed.IsEqual(transposed)
		assert.Equal(t, isCorrect, true)
		assert.Equal(t, err, c.expectedError)
	}

}

func TestMatrixMultiplication(t *testing.T) {
	cases := []struct {
		matrix1        numericalgo.Matrix
		matrix2        numericalgo.Matrix
		expectedResult numericalgo.Matrix
		expectedError  error
	}{
		// Basic
		{
			matrix1: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: numericalgo.Matrix{
				{1, 1},
				{2, 3},
				{5, 2},
			},
			expectedResult: numericalgo.Matrix{
				{20, 13},
				{44, 31},
			},
			expectedError: nil,
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: numericalgo.Matrix{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			expectedResult: numericalgo.Matrix{
				{14, 32},
				{32, 77},
			},
			expectedError: nil,
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: numericalgo.Matrix{
				{1, 4},
				{2, 5},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("The number of columns of the 1st matrix must equal the number of rows of the 2nd matrix"),
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			matrix2: numericalgo.Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expectedResult: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedError: nil,
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: numericalgo.Matrix{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			expectedResult: numericalgo.Matrix{
				{6, 12, 18},
				{15, 30, 45},
			},
			expectedError: nil,
		},
		{
			matrix1: numericalgo.Matrix{
				{3, 4, 2},
			},
			matrix2: numericalgo.Matrix{
				{13, 9, 7, 15},
				{8, 7, 4, 6},
				{6, 4, 0, 3},
			},
			expectedResult: numericalgo.Matrix{
				{83, 63, 37, 75},
			},
			expectedError: nil,
		},
	}

	for _, c := range cases {
		multiplied, err := c.matrix1.MultiplyBy(c.matrix2)
		assert.Equal(t, multiplied, c.expectedResult)
		assert.Equal(t, err, c.expectedError)
	}
}

func TestMatrixLeftDivide(t *testing.T) {
	cases := []struct {
		matrix1        numericalgo.Matrix
		matrix2        numericalgo.Matrix
		expectedResult numericalgo.Matrix
		expectedError  error
	}{
		// Basic
		{
			matrix1: numericalgo.Matrix{
				{2},
				{4},
			},
			matrix2: numericalgo.Matrix{
				{4},
				{4},
			},
			expectedResult: numericalgo.Matrix{
				{1.2},
			},
			expectedError: nil,
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{2, 2},
			},
			matrix2: numericalgo.Matrix{
				{3, 2},
				{1, 1},
			},
			expectedResult: numericalgo.Matrix{
				{-2, -1},
				{2.5, 1.5},
			},
			expectedError: nil,
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{2, 2},
			},
			matrix2: numericalgo.Matrix{
				{3, 2},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("The number of columns of the 1st matrix must equal the number of rows of the 2nd matrix"),
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: numericalgo.Matrix{
				{1, 1},
				{1, 1},
				{1, 1},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Matrix is singular"),
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 1.3},
				{1, 2.1},
				{1, 3.7},
				{1, 4.2},
			},
			matrix2: numericalgo.Matrix{
				{2.2},
				{5.8},
				{10.2},
				{11.8},
			},
			expectedResult: numericalgo.Matrix{
				{-1.5225601452564645},
				{3.1938266000907847},
			},
			expectedError: nil,
		},
		{
			matrix1: numericalgo.Matrix{
				{1, 0.3},
				{1, 0.8},
				{1, 1.2},
				{1, 1.7},
				{1, 2.4},
				{1, 3.1},
				{1, 3.8},
				{1, 4.5},
				{1, 5.1},
				{1, 5.8},
				{1, 6.5},
			},
			matrix2: numericalgo.Matrix{
				{8.61},
				{7.94},
				{7.55},
				{6.85},
				{6.11},
				{5.17},
				{4.19},
				{3.41},
				{2.63},
				{1.77},
				{0.89},
			},
			expectedResult: numericalgo.Matrix{
				{8.99987709451432},
				{-1.246552501126634},
			},
			expectedError: nil,
		},
	}

	for _, c := range cases {
		leftDivided, err := c.matrix1.LeftDivide(c.matrix2)

		isSimilar := leftDivided.IsSimilar(c.expectedResult, 1e-4)
		assert.Equal(t, true, isSimilar)
		assert.Equal(t, err, c.expectedError)
	}
}

func TestMatrixInverse(t *testing.T) {
	cases := []struct {
		matrix         numericalgo.Matrix
		expectedResult numericalgo.Matrix
		expectedError  error
	}{
		{
			matrix: numericalgo.Matrix{
				{4, 7},
				{2, 6},
			},
			expectedResult: numericalgo.Matrix{
				{0.6, -0.7},
				{-0.2, 0.4},
			},
			expectedError: nil,
		},
		{
			matrix: numericalgo.Matrix{
				{4, 7},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Cannot invert non-square Matrix"),
		},
		{
			matrix: numericalgo.Matrix{
				{2, 4},
				{6, 12},
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Matrix is singular"),
		},
		{
			matrix: numericalgo.Matrix{
				{3, 0, 2},
				{2, 0, -2},
				{0, 1, 1},
			},
			expectedResult: numericalgo.Matrix{
				{0.2, 0.2, 0},
				{-0.2, 0.3, 1},
				{0.2, -0.3, 0},
			},
			expectedError: nil,
		},
	}

	for _, c := range cases {
		inverted, err := c.matrix.Invert()
		isSimilar := inverted.IsSimilar(c.expectedResult, 1e-10)
		assert.Equal(t, true, isSimilar)
		assert.Equal(t, err, c.expectedError)
	}
}
func TestMatrixIsSimilar(t *testing.T) {
	cases := []struct {
		matrix1        numericalgo.Matrix
		matrix2        numericalgo.Matrix
		expectedResult bool
	}{
		// Basic
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			matrix2: numericalgo.Matrix{
				{1.000000001, 2.0000000001},
				{3, 4},
			},
			expectedResult: true,
		},
		// Wrong Dimensions
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
			},
			matrix2: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			expectedResult: false,
		},
		// Wrong values
		{
			matrix1: numericalgo.Matrix{
				{1.2, 2.5},
				{3.2, 5.4},
			},
			matrix2: numericalgo.Matrix{
				{1, 2},
				{3, 4},
			},
			expectedResult: false,
		},
		// Nil passed
		{
			matrix1: numericalgo.Matrix{
				{1, 2},
				{3, 5},
			},
			matrix2:        nil,
			expectedResult: false,
		},
	}

	for _, c := range cases {
		isSimilar := c.matrix1.IsSimilar(c.matrix2, 0.1)
		assert.Equal(t, c.expectedResult, isSimilar)
	}
}
