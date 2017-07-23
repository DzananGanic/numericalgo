package numericalgo_test

import (
	"fmt"
	"testing"

	"github.com/DzananGanic/numericalgo"
	"github.com/stretchr/testify/assert"
)

func TestCreateVector(t *testing.T) {
	// Creating empty matrix
	_ = make(numericalgo.Vector, 2)

	// With predefined values
	_ = numericalgo.Vector{1, 2}
}

func TestVectorDim(t *testing.T) {
	cases := []struct {
		vector         numericalgo.Vector
		expectedResult int
	}{
		{
			vector:         numericalgo.Vector{1, 2, 3},
			expectedResult: 3,
		},
		{
			vector:         numericalgo.Vector{},
			expectedResult: 0,
		},
		{
			vector:         numericalgo.Vector{1},
			expectedResult: 1,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedResult, c.vector.Dim())
	}

}

func TestVectorIsDimEqual(t *testing.T) {
	cases := []struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult bool
	}{
		{
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{3, 1, 0},
			expectedResult: true,
		},
		{
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: false,
		},
		{
			vector1:        numericalgo.Vector{1},
			vector2:        numericalgo.Vector{},
			expectedResult: false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedResult, c.vector1.IsDimEqual(c.vector2))
	}

}

func TestAddVectors(t *testing.T) {
	cases := []struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult numericalgo.Vector
		expectedError  error
	}{
		{
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{3, 1, 0},
			expectedResult: numericalgo.Vector{4, 3, 3},
			expectedError:  nil,
		},
		{
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for _, c := range cases {
		result, err := c.vector1.Add(c.vector2)

		assert.Equal(t, c.expectedResult, result)
		assert.Equal(t, c.expectedError, err)
	}

}

func TestSubtractVectors(t *testing.T) {
	cases := []struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult numericalgo.Vector
		expectedError  error
	}{
		{
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{3, 1, 0},
			expectedResult: numericalgo.Vector{-2, 1, 3},
			expectedError:  nil,
		},
		{
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for _, c := range cases {
		result, err := c.vector1.Subtract(c.vector2)

		assert.Equal(t, c.expectedResult, result)
		assert.Equal(t, c.expectedError, err)
	}

}

func TestVectorDotProduct(t *testing.T) {
	cases := []struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult float64
		expectedError  error
	}{
		{
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{4, 5, 6},
			expectedResult: 32,
			expectedError:  nil,
		},
		{
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: 0,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for _, c := range cases {
		result, err := c.vector1.Dot(c.vector2)

		assert.Equal(t, c.expectedResult, result)
		assert.Equal(t, c.expectedError, err)
	}

}

func TestMultiplyVectorByScalar(t *testing.T) {
	cases := []struct {
		vector         numericalgo.Vector
		scalar         float64
		expectedResult numericalgo.Vector
	}{
		{
			vector:         numericalgo.Vector{1, 2, 3},
			scalar:         5,
			expectedResult: numericalgo.Vector{5, 10, 15},
		},
		{
			vector:         numericalgo.Vector{1, 2},
			scalar:         0,
			expectedResult: numericalgo.Vector{0, 0},
		},
	}

	for _, c := range cases {
		result := c.vector.MultiplyByScalar(c.scalar)
		assert.Equal(t, c.expectedResult, result)
	}

}

func TestDivideVectorByScalar(t *testing.T) {
	cases := []struct {
		vector         numericalgo.Vector
		scalar         float64
		expectedResult numericalgo.Vector
		expectedError  error
	}{
		{
			vector:         numericalgo.Vector{1, 2, 3},
			scalar:         5.0,
			expectedResult: numericalgo.Vector{0.2, 0.4, 0.6},
			expectedError:  nil,
		},
		{
			vector:         numericalgo.Vector{1, 2},
			scalar:         0,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Cannot divide by zero"),
		},
	}

	for _, c := range cases {
		result, err := c.vector.DivideByScalar(c.scalar)
		assert.Equal(t, c.expectedResult, result)
		assert.Equal(t, c.expectedError, err)
	}

}
