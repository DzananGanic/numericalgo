package linear_test

import (
	"testing"

	"fmt"

	"github.com/DzananGanic/numericalgo/interpolate"
	"github.com/DzananGanic/numericalgo/interpolate/linear"
	"github.com/stretchr/testify/assert"
)

func TestLinearCanFit(t *testing.T) {
	cases := map[string]struct {
		x             []float64
		y             []float64
		expectedError error
	}{
		"basic linear fit": {
			x:             []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:             []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			expectedError: nil,
		},
		"wrong x and y size": {
			x:             []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 4.07},
			y:             []float64{3.37, 4.45, 4.81, 3.96, 3.31},
			expectedError: fmt.Errorf("X and Y sizes do not match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			li := linear.New()
			err := li.Fit(c.x, c.y)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestLinearCanInterpolateSingleValue(t *testing.T) {
	cases := map[string]struct {
		x                  []float64
		y                  []float64
		valueToInterpolate float64
		expectedEstimate   float64
		expectedError      error
	}{
		"basic linear single-valued interpolation": {
			x:                  []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                  []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valueToInterpolate: 5.1,
			expectedEstimate:   3.1566666666666663,
			expectedError:      nil,
		},
		"basic linear single-valued interpolation - middle detection err": {
			x:                  []float64{1100, 1200, 1230, 1260, 1280, 1300, 1320, 1340, 1380, 1440, 1590},
			y:                  []float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			valueToInterpolate: 1265,
			expectedEstimate:   32.5,
			expectedError:      nil,
		},
		"testing binary search for nearest neighbor - case where the interpolation value should be between indexes 0 and 1": {
			x:                  []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                  []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valueToInterpolate: 1.5,
			expectedEstimate:   3.802,
			expectedError:      nil,
		},
		"testing binary search for nearest neighbor - case where the interpolation value should be between last two indexes": {
			x:                  []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                  []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valueToInterpolate: 5.8,
			expectedEstimate:   3.704285714285714,
			expectedError:      nil,
		},
		"unsorted x and y test": {
			x:                  []float64{1.8, 4.9, 2.5, 1.3, 4.4, 3.1, 3.8, 5.5, 6.2},
			y:                  []float64{4.45, 3.02, 4.81, 3.37, 2.72, 3.96, 3.31, 3.43, 4.07},
			valueToInterpolate: 2.2,
			expectedEstimate:   4.655714285714286,
			expectedError:      nil,
		},
		"big value to interpolate test": {
			x:                  []float64{1.8, 4.9, 2.5, 1.3, 4.4, 3.1, 3.8, 5.5, 6.2},
			y:                  []float64{4.45, 3.02, 4.81, 3.37, 2.72, 3.96, 3.31, 3.43, 4.07},
			valueToInterpolate: 1000,
			expectedEstimate:   0,
			expectedError:      fmt.Errorf("Value to interpolate is too large and not in range"),
		},
		"too small value to interpolate test": {
			x:                  []float64{1.8, 4.9, 2.5, 1.3, 4.4, 3.1, 3.8, 5.5, 6.2},
			y:                  []float64{4.45, 3.02, 4.81, 3.37, 2.72, 3.96, 3.31, 3.43, 4.07},
			valueToInterpolate: -20,
			expectedEstimate:   0,
			expectedError:      fmt.Errorf("Value to interpolate is too small and not in range"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			li := linear.New()
			li.Fit(c.x, c.y)
			estimate, err := interpolate.WithSingle(li, c.valueToInterpolate)
			assert.Equal(t, c.expectedEstimate, estimate)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestLinearCanInterpolateMultipleValues(t *testing.T) {
	cases := map[string]struct {
		x                   []float64
		y                   []float64
		valuesToInterpolate []float64
		expectedEstimates   []float64
		expectedError       error
	}{
		"basic linear multiple-value interpolation": {
			x:                   []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                   []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valuesToInterpolate: []float64{2.2, 5.1, 1.5},
			expectedEstimates:   []float64{4.655714285714286, 3.1566666666666663, 3.802},
			expectedError:       nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			li := linear.New()
			li.Fit(c.x, c.y)
			estimates, err := interpolate.WithMulti(li, c.valuesToInterpolate)
			assert.Equal(t, c.expectedEstimates, estimates)
			assert.Equal(t, c.expectedError, err)
		})
	}
}
