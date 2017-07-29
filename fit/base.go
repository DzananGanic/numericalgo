package fit

import (
	"github.com/DzananGanic/numericalgo"
)

// Base is a basic struct type on which other fit types are built on. It has x, y and coeff vectors as properties, as well as methods for getting them.
type Base struct {
	x     numericalgo.Vector
	y     numericalgo.Vector
	coeff numericalgo.Vector
}

// X returns the slice of the fitted X coordinates
func (b *Base) X() []float64 {
	return b.x
}

// Y returns the slice of the fitted Y coordinates
func (b *Base) Y() []float64 {
	return b.y
}

// Coeff returns the slice of the fit coefficients
func (b *Base) Coeff() numericalgo.Vector {
	return b.coeff
}
