package exponential_test

import (
	"testing"

	"github.com/teivah/numericalgo"
	"github.com/teivah/numericalgo/fit"
	"github.com/teivah/numericalgo/fit/exponential"
	"github.com/stretchr/testify/assert"
)

func TestFitExponentialFit(t *testing.T) {
	cases := map[string]struct {
		x             numericalgo.Vector
		y             numericalgo.Vector
		p             float64
		q             float64
		expectedError error
	}{
		"basic exponential fit": {
			x:             numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:             numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			p:             6.9758664327105,
			q:             0.5471581141119718,
			expectedError: nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ef := exponential.New()
			err := ef.Fit(c.x, c.y)
			coef := ef.Coeff
			assert.InEpsilon(t, coef[0], c.p, 1e-10)
			assert.InEpsilon(t, coef[1], c.q, 1e-10)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestPredictExponentialFit(t *testing.T) {

	cases := map[string]struct {
		x              numericalgo.Vector
		y              numericalgo.Vector
		valueToPredict float64
		expectedResult float64
		expectedError  error
	}{
		"basic predict exponential fit": {
			x:              numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:              numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			valueToPredict: 1.2,
			expectedResult: 7.632456,
			expectedError:  nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ef := exponential.New()
			err := ef.Fit(c.x, c.y)
			result := ef.Predict(c.valueToPredict)
			assert.InEpsilon(t, c.expectedResult, result, 1e-4)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestPredictMultiExponentialFit(t *testing.T) {
	cases := map[string]struct {
		x               numericalgo.Vector
		y               numericalgo.Vector
		valuesToPredict numericalgo.Vector
		expectedResult  numericalgo.Vector
		expectedError   error
	}{
		"basic multi exponential fit prediction": {
			x:               numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:               numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			valuesToPredict: numericalgo.Vector{1.9, 7.0, 8.0},
			expectedResult:  numericalgo.Vector{8.01546, 10.8059732, 11.353131},
			expectedError:   nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ef := exponential.New()
			err := ef.Fit(c.x, c.y)
			result := fit.PredictMulti(ef, c.valuesToPredict)
			assert.InEpsilonSlice(t, c.expectedResult, result, 1e-4)
			assert.Equal(t, c.expectedError, err)
		})
	}
}
