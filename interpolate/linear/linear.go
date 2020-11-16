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

func (li *Linear) Interpolate(val float64) float64 {
	var est float64

	l, r := li.findNearestNeighbors(val, 0, len(li.XYPairs)-1)

	lX := li.XYPairs[l].X
	rX := li.XYPairs[r].X
	lY := li.XYPairs[l].Y
	rY := li.XYPairs[r].Y

	est = lY + (rY-lY)/(rX-lX)*(val-lX)
	return est
}

func (li *Linear) Validate(val float64) error {

	if val < li.XYPairs[0].X {
		return fmt.Errorf("Value to interpolate is too small and not in range")
	}

	if val > li.XYPairs[len(li.XYPairs)-1].X {
		return fmt.Errorf("Value to interpolate is too large and not in range")
	}

	return nil
}

func (li *Linear) findNearestNeighbors(val float64, l, r int) (int, int) {
	middle := (l + r) / 2
	if (val >= li.XYPairs[middle-1].X) && (val <= li.XYPairs[middle].X) {
		return middle - 1, middle
	} else if val < li.XYPairs[middle-1].X {
		return li.findNearestNeighbors(val, l, middle-1)
	}
	return li.findNearestNeighbors(val, middle+1, r)
}
