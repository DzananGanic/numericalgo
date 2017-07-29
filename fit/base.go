package fit

import (
	"github.com/DzananGanic/numericalgo"
)

type Base struct {
	x     []float64
	y     []float64
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

func (b *Base) Coef() numericalgo.Vector {
	return b.coeff
}
