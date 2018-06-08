package root_test

import (
	"math"
	"testing"

	"github.com/teivah/numericalgo/root"
	"github.com/stretchr/testify/assert"
)

func TestNewton(t *testing.T) {

	cases := map[string]struct {
		f             func(x float64) float64
		iter          int
		initialGuess  float64
		expectedValue float64
		expectedError error
	}{
		"basic root finding": {
			f: func(x float64) float64 {
				return math.Pow(x, 3) - 2*math.Pow(x, 2) + 5
			},
			iter:          3,
			initialGuess:  -1,
			expectedValue: -1.2419,
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := root.Newton(c.f, c.initialGuess, c.iter)
			assert.InEpsilon(t, result, c.expectedValue, 1e-4)
			assert.Equal(t, err, c.expectedError)
		})
	}
}
