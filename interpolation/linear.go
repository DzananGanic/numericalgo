package interpolation

import "fmt"

// Linear provides the basic functionality for linear interpolation.
// Given X and Y float64 slices, it can estimate the value of the function at the desired point.
type Linear struct {
	Base
}

// NewLinear returns the new Linear object
func NewLinear() *Linear {
	li := &Linear{}
	return li
}

// Interpolate receives valueToInterpolate float64 parameter, and returns the estimate by using linear interpolation formula.
func (li *Linear) Interpolate(valueToInterpolate float64) float64 {
	var estimate float64

	left, right := li.findNearestNeighbors(valueToInterpolate, 0, len(li.xyPairs)-1)

	leftX := li.xyPairs[left].X
	rightX := li.xyPairs[right].X
	leftY := li.xyPairs[left].Y
	rightY := li.xyPairs[right].Y
	estimate = leftY + (rightY-leftY)/(rightX-leftX)*(valueToInterpolate-leftX)

	return estimate
}

func (li *Linear) findNearestNeighbors(valueToInterpolate float64, left, right int) (int, int) {
	middle := (left + right) / 2
	if (valueToInterpolate >= li.xyPairs[middle-1].X) && (valueToInterpolate <= li.xyPairs[middle].X) {
		return middle - 1, middle
	} else if valueToInterpolate < li.xyPairs[middle-1].X {
		return li.findNearestNeighbors(valueToInterpolate, left, middle-2)
	}
	return li.findNearestNeighbors(valueToInterpolate, middle+1, right)
}

func (li *Linear) validate(valueToInterpolate float64) error {

	if valueToInterpolate < li.xyPairs[0].X {
		return fmt.Errorf("Value to interpolate is too small and not in range")
	}

	if valueToInterpolate > li.xyPairs[len(li.xyPairs)-1].X {
		return fmt.Errorf("Value to interpolate is too large and not in range")
	}

	return nil
}
