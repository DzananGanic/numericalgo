package linear

import (
	"fmt"

	"github.com/DzananGanic/numericalgo/interpolate"
)

// Linear provides the basic functionality for linear interpolation.
// Given X and Y float64 slices, it can estimate the value of the function at the desired point.
type Linear struct {
	interpolate.Base
}

// New returns the new Linear object
func New() *Linear {
	li := &Linear{}
	return li
}

func (li *Linear) Interpolate(valueToInterpolate float64) float64 {
	var estimate float64

	left, right := li.findNearestNeighbors(valueToInterpolate, 0, len(li.XYPairs)-1)

	leftX := li.XYPairs[left].X
	rightX := li.XYPairs[right].X
	leftY := li.XYPairs[left].Y
	rightY := li.XYPairs[right].Y
	estimate = leftY + (rightY-leftY)/(rightX-leftX)*(valueToInterpolate-leftX)

	return estimate
}

func (li *Linear) Validate(valueToInterpolate float64) error {

	if valueToInterpolate < li.XYPairs[0].X {
		return fmt.Errorf("Value to interpolate is too small and not in range")
	}

	if valueToInterpolate > li.XYPairs[len(li.XYPairs)-1].X {
		return fmt.Errorf("Value to interpolate is too large and not in range")
	}

	return nil
}

func (li *Linear) findNearestNeighbors(valueToInterpolate float64, left, right int) (int, int) {
	middle := (left + right) / 2
	if (valueToInterpolate >= li.XYPairs[middle-1].X) && (valueToInterpolate <= li.XYPairs[middle].X) {
		return middle - 1, middle
	} else if valueToInterpolate < li.XYPairs[middle-1].X {
		return li.findNearestNeighbors(valueToInterpolate, left, middle-2)
	}
	return li.findNearestNeighbors(valueToInterpolate, middle+1, right)
}
