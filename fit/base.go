package fit

type Base struct {
	x []float64
	y []float64
	p float64
	q float64
}

// X returns the slice of the fitted X coordinates
func (b *Base) X() []float64 {
	return b.x
}

// Y returns the slice of the fitted Y coordinates
func (b *Base) Y() []float64 {
	return b.y
}

func (b *Base) Coef() (float64, float64) {
	return b.p, b.q
}

func (b *Base) Predict(val float64) float64 {
	return b.p*val + b.q
}
