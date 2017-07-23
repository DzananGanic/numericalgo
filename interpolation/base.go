package interpolation

import (
	"fmt"

	"github.com/DzananGanic/numericalgo"
)

// Base type provides the base functionality for any interpolation type
type Base struct {
	xyPairs []numericalgo.CoordinatePair
	x       []float64
	y       []float64
}

// X returns the slice of the fitted X coordinates
func (b *Base) X() []float64 {
	return b.x
}

// Y returns the slice of the fitted Y coordinates
func (b *Base) Y() []float64 {
	return b.y
}

// XYPairs returns the slice of the fitted XY coordinate pairs
func (b *Base) XYPairs() []numericalgo.CoordinatePair {
	return b.xyPairs
}

// FitSamples receives two float64 slices - for the x and y coordinates where x[i] and y[i] represent a coordinate pair in a grid.
// It returns the error if the X and Y sizes do not match.
func (b *Base) FitSamples(x, y []float64) error {
	if len(x) != len(y) {
		return fmt.Errorf("X and Y sizes do not match")
	}
	b.x = x
	b.y = y
	b.xyPairs = numericalgo.SlicesToCoordinatePairs(x, y)
	numericalgo.SortCoordinatePairs(b.xyPairs)
	return nil
}
