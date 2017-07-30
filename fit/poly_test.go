package fit_test

import (
	"testing"

	"github.com/DzananGanic/numericalgo"
	"github.com/DzananGanic/numericalgo/fit"
	"github.com/stretchr/testify/assert"
)

func TestFitPolyFit(t *testing.T) {
	cases := map[string]struct {
		x             numericalgo.Vector
		y             numericalgo.Vector
		n             int
		coef          numericalgo.Vector
		expectedError error
	}{
		"basic poly fit": {
			x:             numericalgo.Vector{0.0, 1.0, 2.0, 3.0, 4.0, 5.0},
			y:             numericalgo.Vector{0.0, 0.8, 0.9, 0.1, -0.8, -1.0},
			n:             3,
			coef:          numericalgo.Vector{-0.0396825, 1.693121, -0.8134920, 0.0870370},
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			pf := fit.NewPoly()
			err := pf.Fit(c.x, c.y, c.n)
			r := pf.Coeff().IsSimilar(c.coef, 1e-4)
			assert.Equal(t, true, r)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestPredictPolyFit(t *testing.T) {
	cases := map[string]struct {
		x              []float64
		y              []float64
		n              int
		valueToPredict float64
		expectedResult float64
		expectedError  error
	}{
		"basic poly fit prediction": {
			x:              numericalgo.Vector{0.0, 1.0, 2.0, 3.0, 4.0, 5.0},
			y:              numericalgo.Vector{0.0, 0.8, 0.9, 0.1, -0.8, -1.0},
			n:              3,
			valueToPredict: 2.0,
			expectedResult: 0.7888888888889196,
			expectedError:  nil,
		},
		"second poly fit prediction": {
			x:              numericalgo.Vector{0.0, 1.0, 2.0, 3.0, 4.0, 5.0},
			y:              numericalgo.Vector{0.0, 0.8, 0.9, 0.1, -0.8, -1.0},
			n:              3,
			valueToPredict: 2.5,
			expectedResult: 0.4687499999999935,
			expectedError:  nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			pf := fit.NewPoly()
			err := pf.Fit(c.x, c.y, c.n)
			result := pf.Predict(c.valueToPredict)
			assert.InEpsilon(t, c.expectedResult, result, 1e-2)
			assert.Equal(t, c.expectedError, err)
		})
	}
}
