package fit_test

import (
	"fmt"
	"testing"

	"github.com/DzananGanic/numericalgo"
	"github.com/DzananGanic/numericalgo/fit"
	"github.com/stretchr/testify/assert"
)

func TestFitExponentialFit(t *testing.T) {
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
			p:             1311.0917974555966,
			q:             0.07892996161919555,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		ef := fit.NewExponential()
		err := ef.Fit(c.x, c.y)
		coef := ef.Coeff()
		assert.InEpsilon(t, coef[0], c.p, 1e-10)
		assert.InEpsilon(t, coef[1], c.q, 1e-10)
		assert.Equal(t, c.expectedError, err)
	}
}

func TestPredictExponentialFit(t *testing.T) {

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
			valueToPredict: 1.2,
			expectedResult: 6.631427342373716,
			expectedError:  nil,
		},
	}

	for _, c := range cases {
		ef := fit.NewExponential()
		err := ef.Fit(c.x, c.y)
		result := ef.Predict(c.valueToPredict)
		fmt.Println(result)
		assert.InEpsilon(t, c.expectedResult, result, 1e-4)
		assert.Equal(t, c.expectedError, err)
	}
}
