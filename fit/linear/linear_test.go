package linear_test

import (
	"testing"

	"github.com/teivah/numericalgo"
	"github.com/teivah/numericalgo/fit"
	"github.com/teivah/numericalgo/fit/linear"
	"github.com/stretchr/testify/assert"
)

func TestFitLinearFit(t *testing.T) {
	cases := map[string]struct {
		x             numericalgo.Vector
		y             numericalgo.Vector
		p             float64
		q             float64
		expectedError error
	}{
		"basic linear fit": {
			x:             numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:             numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			p:             -1.246552501126634,
			q:             8.99987709451432,
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			lf := linear.New()
			err := lf.Fit(c.x, c.y)
			coef := lf.Coeff
			assert.InEpsilon(t, coef[1], c.p, 1e-10)
			assert.InEpsilon(t, coef[0], c.q, 1e-10)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestPredictLinearFit(t *testing.T) {

	cases := map[string]struct {
		x              numericalgo.Vector
		y              numericalgo.Vector
		valueToPredict float64
		expectedResult float64
		expectedError  error
	}{
		"basic linear fit prediction": {
			x:              numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:              numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			valueToPredict: 1.9,
			expectedResult: 6.631427342373716,
			expectedError:  nil,
		},
		"second linear fit prediction": {
			x:              numericalgo.Vector{1.3, 2.1, 3.7, 4.2},
			y:              numericalgo.Vector{2.2, 5.8, 10.2, 11.8},
			valueToPredict: 1.9,
			expectedResult: 3.19383*1.9 - 1.52256,
			expectedError:  nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			lf := linear.New()
			err := lf.Fit(c.x, c.y)
			result := lf.Predict(c.valueToPredict)
			assert.InEpsilon(t, c.expectedResult, result, 1e-4)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestPredictMultiLinearFit(t *testing.T) {
	cases := map[string]struct {
		x               numericalgo.Vector
		y               numericalgo.Vector
		valuesToPredict numericalgo.Vector
		expectedResult  numericalgo.Vector
		expectedError   error
	}{
		"basic multi linear fit prediction": {
			x:               numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:               numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			valuesToPredict: numericalgo.Vector{1.9, 7.0, 8.0},
			expectedResult:  numericalgo.Vector{6.631427, 0.27400958, -0.972542},
			expectedError:   nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			lf := linear.New()
			err := lf.Fit(c.x, c.y)
			result := fit.PredictMulti(lf, c.valuesToPredict)
			assert.InEpsilonSlice(t, c.expectedResult, result, 1e-4)
			assert.Equal(t, c.expectedError, err)
		})
	}
}
