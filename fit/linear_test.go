package fit_test

import (
	"testing"

	"github.com/DzananGanic/numericalgo"
	"github.com/DzananGanic/numericalgo/fit"
	"github.com/stretchr/testify/assert"
)

func TestCreateLinearFit(t *testing.T) {
	lf := fit.NewLinear()
	_ = lf
}

func TestFitLinearFit(t *testing.T) {
	cases := []struct {
		x             numericalgo.Vector
		y             numericalgo.Vector
		p             float64
		q             float64
		expectedError error
	}{
		//Basic
		{
			x:             numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:             numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			p:             -1.246552501126634,
			q:             8.99987709451432,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		lf := fit.NewLinear()
		err := lf.Fit(c.x, c.y)
		p, q := lf.Coef()
		assert.InEpsilon(t, p, c.p, 1e-10)
		assert.InEpsilon(t, q, c.q, 1e-10)
		assert.Equal(t, c.expectedError, err)
	}
}

func TestPredictLinearFit(t *testing.T) {

	cases := []struct {
		x              []float64
		y              []float64
		valueToPredict float64
		expectedResult float64
		expectedError  error
	}{
		//Basic
		{
			x:              []float64{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5},
			y:              []float64{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89},
			valueToPredict: 1.9,
			expectedResult: 6.631427342373716,
			expectedError:  nil,
		},
		{
			x:              []float64{1.3, 2.1, 3.7, 4.2},
			y:              []float64{2.2, 5.8, 10.2, 11.8},
			valueToPredict: 1.9,
			expectedResult: 3.19383*1.9 - 1.52256,
			expectedError:  nil,
		},
	}

	for _, c := range cases {
		lf := fit.NewLinear()
		err := lf.Fit(c.x, c.y)
		result := lf.Predict(c.valueToPredict)
		assert.InEpsilon(t, c.expectedResult, result, 1e-4)
		assert.Equal(t, c.expectedError, err)
	}
}
