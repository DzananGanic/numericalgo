package lagrange_test

import (
	"fmt"
	"testing"

	"github.com/DzananGanic/numericalgo/interpolate"
	"github.com/DzananGanic/numericalgo/interpolate/lagrange"
	"github.com/stretchr/testify/assert"
)

func TestLagrangeCanFit(t *testing.T) {
	cases := map[string]struct {
		x             []float64
		y             []float64
		expectedError error
	}{
		"basic lagrange fit": {
			x:             []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:             []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			expectedError: nil,
		},
		"wrong x and y sizes lagrange fit": {
			x:             []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 4.07},
			y:             []float64{3.37, 4.45, 4.81, 3.96, 3.31},
			expectedError: fmt.Errorf("X and Y sizes do not match"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			lgi := lagrange.New()
			//lgi := interpolation.NewLagrange()
			err := lgi.Fit(c.x, c.y)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestLagrangeCanInterpolateSingleValue(t *testing.T) {
	cases := map[string]struct {
		x                  []float64
		y                  []float64
		valueToInterpolate float64
		expectedEstimate   float64
		expectedError      error
	}{
		"basic lagrange single-value interpolation": {
			x:                  []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                  []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valueToInterpolate: 5.1,
			expectedEstimate:   3.3068917458526563,
			expectedError:      nil,
		},
		"testing binary search for nearest neighbor - case where the interpolation value should be between indexes 0 and 1": {
			x:                  []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                  []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valueToInterpolate: 1.5,
			expectedEstimate:   3.224674773993458,
			expectedError:      nil,
		},
		"testing binary search for nearest neighbor - case where the interpolation value should be between last two indexes": {
			x:                  []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                  []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valueToInterpolate: 5.8,
			expectedEstimate:   2.785117811403902,
			expectedError:      nil,
		},
		"unsorted x and y test": {
			x:                  []float64{1.8, 4.9, 2.5, 1.3, 4.4, 3.1, 3.8, 5.5, 6.2},
			y:                  []float64{4.45, 3.02, 4.81, 3.37, 2.72, 3.96, 3.31, 3.43, 4.07},
			valueToInterpolate: 2.2,
			expectedEstimate:   5.134038366257703,
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
		"same x values error": {
			x:                  []float64{1.8, 1.8, 1.8, 1.3, 4.4, 3.1, 3.8, 5.5, 6.2},
			y:                  []float64{4.45, 3.02, 4.81, 3.37, 2.72, 3.96, 3.31, 3.43, 4.07},
			valueToInterpolate: -20,
			expectedEstimate:   0,
			expectedError:      fmt.Errorf("There are at least 2 same X values. This will result in division by zero in Lagrange interpolation"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			lg := lagrange.New()
			lg.Fit(c.x, c.y)
			estimate, err := interpolate.WithSingle(lg, c.valueToInterpolate)
			assert.Equal(t, c.expectedEstimate, estimate)
			assert.Equal(t, c.expectedError, err)
		})
	}
}

func TestLagrangeCanInterpolateMultipleValues(t *testing.T) {
	cases := map[string]struct {
		x                   []float64
		y                   []float64
		valuesToInterpolate []float64
		expectedEstimates   []float64
		expectedError       error
	}{
		"basic lagrange multiple-value interpolation": {
			x:                   []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2},
			y:                   []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07},
			valuesToInterpolate: []float64{2.2, 5.1, 1.5},
			expectedEstimates:   []float64{5.134038366257704, 3.3068917458526563, 3.224674773993458},
			expectedError:       nil,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			lg := lagrange.New()
			lg.Fit(c.x, c.y)
			estimates, err := interpolate.WithMulti(lg, c.valuesToInterpolate)
			assert.Equal(t, c.expectedEstimates, estimates)
			assert.Equal(t, c.expectedError, err)
		})
	}
}
