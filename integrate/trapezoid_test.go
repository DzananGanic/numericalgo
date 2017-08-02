package integrate_test

import (
	"math"
	"testing"

	"github.com/DzananGanic/numericalgo/integrate"
	"github.com/stretchr/testify/assert"
)

func TestTrapezoid(t *testing.T) {

	cases := map[string]struct {
		f             func(x float64) float64
		l             float64
		r             float64
		n             int
		expectedValue float64
		expectedError error
	}{
		"f(x) = x^4 with n = 20": {
			f: func(x float64) float64 {
				return math.Pow(x, 4)
			},
			l:             1,
			r:             3,
			n:             20,
			expectedValue: 48.48666,
			expectedError: nil,
		},
		"f(x) = x^4 with n = 50": {
			f: func(x float64) float64 {
				return math.Pow(x, 4)
			},
			l:             1,
			r:             3,
			n:             50,
			expectedValue: 48.41386,
			expectedError: nil,
		},
		"f(x) = x^4 with n = 100": {
			f: func(x float64) float64 {
				return math.Pow(x, 4)
			},
			l:             1,
			r:             3,
			n:             100,
			expectedValue: 48.400138,
			expectedError: nil,
		},
		"f(x) = x^4 with n = 500": {
			f: func(x float64) float64 {
				return math.Pow(x, 4)
			},
			l:             1,
			r:             3,
			n:             100,
			expectedValue: 48.400138,
			expectedError: nil,
		},
		"f(x) = 1/x with n = 20": {
			f: func(x float64) float64 {
				return 1 / x
			},
			l:             2,
			r:             5,
			n:             20,
			expectedValue: 0.91668422,
			expectedError: nil,
		},
		"f(x) = sin(x) with n = 20": {
			f: func(x float64) float64 {
				return math.Sin(x)
			},
			l:             0,
			r:             math.Pi / 2,
			n:             20,
			expectedValue: 0.9994859,
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := integrate.Trapezoid(c.f, c.l, c.r, c.n)
			assert.InEpsilon(t, result, c.expectedValue, 1e-4)
			assert.Equal(t, err, c.expectedError)
		})
	}
}
