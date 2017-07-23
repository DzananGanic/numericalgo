package interpolation

import "fmt"

// Lagrange provides the basic functionality for lagrange interpolation.
// Given X and Y float64 slices, it can estimate the value of the function at the desired point.
type Lagrange struct {
	Base
}

// NewLagrange returns the new Lagrange object.
func NewLagrange() *Lagrange {
	lg := &Lagrange{}
	return lg
}

// Interpolate receives valueToInterpolate float64 parameter, and returns the estimate by using lagrange interpolation formula.
func (lg *Lagrange) Interpolate(valueToInterpolate float64) float64 {
	var estimate float64

	for i := 0; i < len(lg.x); i++ {
		product := lg.y[i]
		for j := 0; j < len(lg.x); j++ {
			if i != j {
				product = product * (valueToInterpolate - lg.x[j]) / (lg.x[i] - lg.x[j])
			}
		}
		estimate += product
	}

	return estimate
}

// validate receives valueToInterpolate float64 parameter, and returns the error in input if it exists.
func (lg *Lagrange) validate(valueToInterpolate float64) error {

	// TODO: Check case where lg.x[i]-lg.x[j] is 0

	if valueToInterpolate < lg.xyPairs[0].X {
		return fmt.Errorf("Value to interpolate is too small and not in range")
	}

	if valueToInterpolate > lg.xyPairs[len(lg.xyPairs)-1].X {
		return fmt.Errorf("Value to interpolate is too large and not in range")
	}

	return nil
}
