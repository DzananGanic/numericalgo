package root_test

import (
	"math"
	"testing"

	"github.com/DzananGanic/numericalgo/root"
	"github.com/stretchr/testify/assert"
)

func TestBisection(t *testing.T) {

	cases := map[string]struct {
		f             func(x float64) float64
		eps           float64
		l             float64
		r             float64
		expectedValue float64
		expectedError error
	}{
		"basic bisection": {
			f: func(x float64) float64 {
				return math.Pow(x, 2)
			},
			eps:           0.01,
			l:             -1,
			r:             1,
			expectedValue: 0,
			expectedError: nil,
		},
		"bisecting more complex function": {
			f: func(x float64) float64 {
				return math.Pow(x, 3) - 2*math.Pow(x, 2) + 5
			},
			eps:           0.01,
			l:             -1.5,
			r:             -1,
			expectedValue: -1.2421875,
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := root.Bisection(c.f, c.eps, c.l, c.r)
			assert.Equal(t, result, c.expectedValue)
			assert.Equal(t, err, c.expectedError)
		})
	}
}
