package numericalgo_test

import (
	"fmt"
	"testing"

	"github.com/teivah/numericalgo"
	"github.com/stretchr/testify/assert"
)

func TestVectorDim(t *testing.T) {
	cases := map[string]struct {
		vector         numericalgo.Vector
		expectedResult int
	}{
		"basic test": {
			vector:         numericalgo.Vector{1, 2, 3},
			expectedResult: 3,
		},
		"empty vector": {
			vector:         numericalgo.Vector{},
			expectedResult: 0,
		},
		"single element vector": {
			vector:         numericalgo.Vector{1},
			expectedResult: 1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, c.expectedResult, c.vector.Dim())
		})
	}

}

func TestVectorAreDimsEqual(t *testing.T) {
	cases := map[string]struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult bool
	}{
		"basic equality test": {
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{3, 1, 0},
			expectedResult: true,
		},
		"different dimensions": {
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: false,
		},
		"testing with empty vector": {
			vector1:        numericalgo.Vector{1},
			vector2:        numericalgo.Vector{},
			expectedResult: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, c.expectedResult, c.vector1.AreDimsEqual(c.vector2))
		})
	}
}

func TestAddVectors(t *testing.T) {
	cases := map[string]struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult numericalgo.Vector
		expectedError  error
	}{
		"add vectors basic test": {
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{3, 1, 0},
			expectedResult: numericalgo.Vector{4, 3, 3},
			expectedError:  nil,
		},
		"add with wrong dimensions": {
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector1.Add(c.vector2)

			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestSubtractVectors(t *testing.T) {
	cases := map[string]struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult numericalgo.Vector
		expectedError  error
	}{
		"subtract vectors basic test": {
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{3, 1, 0},
			expectedResult: numericalgo.Vector{-2, 1, 3},
			expectedError:  nil,
		},
		"subtract with wrong dimensions": {
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: nil,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector1.Subtract(c.vector2)

			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestVectorDotProduct(t *testing.T) {
	cases := map[string]struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		expectedResult float64
		expectedError  error
	}{
		"basic dot product test": {
			vector1:        numericalgo.Vector{1, 2, 3},
			vector2:        numericalgo.Vector{4, 5, 6},
			expectedResult: 32,
			expectedError:  nil,
		},
		"dot product with wrong dimensions": {
			vector1:        numericalgo.Vector{1, 2},
			vector2:        numericalgo.Vector{1, 2, 3},
			expectedResult: 0,
			expectedError:  fmt.Errorf("Dimensions must match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector1.Dot(c.vector2)

			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestMultiplyVectorByScalar(t *testing.T) {
	cases := map[string]struct {
		vector         numericalgo.Vector
		scalar         float64
		expectedResult numericalgo.Vector
	}{
		"basic multiply vector by scalar": {
			vector:         numericalgo.Vector{1, 2, 3},
			scalar:         5,
			expectedResult: numericalgo.Vector{5, 10, 15},
		},
		"multiply vector by 0": {
			vector:         numericalgo.Vector{1, 2},
			scalar:         0,
			expectedResult: numericalgo.Vector{0, 0},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := c.vector.MultiplyByScalar(c.scalar)
			assert.Equal(t, c.expectedResult, result)
		})
	}

}

func TestDivideVectorByScalar(t *testing.T) {
	cases := map[string]struct {
		vector         numericalgo.Vector
		scalar         float64
		expectedResult numericalgo.Vector
		expectedError  error
	}{
		"basic vector division by scalar": {
			vector:         numericalgo.Vector{1, 2, 3},
			scalar:         5.0,
			expectedResult: numericalgo.Vector{0.2, 0.4, 0.6},
			expectedError:  nil,
		},
		"vector division by 0": {
			vector:         numericalgo.Vector{1, 2},
			scalar:         0,
			expectedResult: nil,
			expectedError:  fmt.Errorf("Cannot divide by zero"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := c.vector.DivideByScalar(c.scalar)
			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestVectorPower(t *testing.T) {
	cases := map[string]struct {
		vector         numericalgo.Vector
		power          float64
		expectedResult numericalgo.Vector
	}{
		"vector elements squared": {
			vector:         numericalgo.Vector{1, 2, 3},
			power:          2,
			expectedResult: numericalgo.Vector{1, 4, 9},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := c.vector.Power(c.power)
			assert.Equal(t, c.expectedResult, result)
		})
	}
}

func TestVectorIsSimilar(t *testing.T) {
	cases := map[string]struct {
		vector1        numericalgo.Vector
		vector2        numericalgo.Vector
		tolerance      float64
		expectedResult bool
	}{
		"test non-similar vectors": {
			vector1:        numericalgo.Vector{1.2, 2.5},
			vector2:        numericalgo.Vector{1, 2},
			tolerance:      0.01,
			expectedResult: false,
		},
		"test similar vectors": {
			vector1:        numericalgo.Vector{1.000000001, 2.0000000001},
			vector2:        numericalgo.Vector{1, 2},
			tolerance:      0.1,
			expectedResult: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := c.vector1.IsSimilar(c.vector2, c.tolerance)
			assert.Equal(t, c.expectedResult, result)
		})
	}
}
