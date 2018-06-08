package differentiate_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/teivah/numericalgo/differentiate"
	"github.com/stretchr/testify/assert"
)

func TestCentral(t *testing.T) {

	cases := map[string]struct {
		f             func(x float64) float64
		val           float64
		h             float64
		expectedValue float64
		expectedError error
	}{
		"central difference with 0.1 step size": {
			f: func(x float64) float64 {
				return math.Cos(math.Pow(x, 2) - 2)
			},
			val:           1,
			h:             0.1,
			expectedValue: 1.6609,
			expectedError: nil,
		},
		"central difference with 0.01 step size": {
			f: func(x float64) float64 {
				return math.Cos(math.Pow(x, 2) - 2)
			},
			val:           1,
			h:             0.01,
			expectedValue: 1.6827,
			expectedError: nil,
		},
		"central difference with 0.001 step size": {
			f: func(x float64) float64 {
				return math.Cos(math.Pow(x, 2) - 2)
			},
			val:           1,
			h:             0.001,
			expectedValue: 1.6829,
			expectedError: nil,
		},
		"central difference with wrong step size": {
			f: func(x float64) float64 {
				return math.Cos(math.Pow(x, 2) - 2)
			},
			val:           1,
			h:             -2,
			expectedValue: 0,
			expectedError: fmt.Errorf("Step size has to be greater than 0"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := differentiate.Central(c.f, c.val, c.h)
			if result != 0 {
				assert.InEpsilon(t, result, c.expectedValue, 1e-4)
			} else {
				assert.Equal(t, result, c.expectedValue)
			}
			assert.Equal(t, err, c.expectedError)
		})
	}
}
