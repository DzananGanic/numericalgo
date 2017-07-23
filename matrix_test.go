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
		{
			matrix: numericalgo.Matrix{
				{1, 4},
				{2},
			},
			transposed:    nil,
			expectedError: fmt.Errorf("Inconsistent dimensions"),
		},
	}

	for _, c := range cases {
		transposed, err := c.matrix.Transpose()
		isCorrect := c.transposed.IsEqual(transposed)
		assert.Equal(t, isCorrect, true)
		assert.Equal(t, err, c.expectedError)
	}

}
