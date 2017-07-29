package fit

import (
	"math"

	"github.com/DzananGanic/numericalgo"
)

type Poly struct {
	Base
}

func NewPoly() *Poly {
	pf := &Poly{}
	return pf
}

func (p *Poly) Fit(x numericalgo.Vector, y numericalgo.Vector, n int) error {
	xMatrix := numericalgo.Matrix{x}
	yMatrix := numericalgo.Matrix{y}

	xT, err := xMatrix.Transpose()

	if err != nil {
		return err
	}

	ones := make(numericalgo.Vector, x.Dim())
	for i := range ones {
		ones[i] = 1
	}

	X, err := xT.AddColumnAt(0, ones)

	if err != nil {
		return err
	}

	for i := 2; i <= n; i++ {
		X, err = X.AddColumnAt(i, x.Power(float64(i)))
	}

	Y, err := yMatrix.Transpose()

	if err != nil {
		return err
	}

	coeff, err := X.LeftDivide(Y)

	if err != nil {
		return err
	}

	p.coeff, err = coeff.GetColumnAt(0)

	if err != nil {
		return err
	}

	return nil
}

func (p *Poly) Predict(val float64) float64 {
	var result float64
	c := p.Coef()
	for i := len(c) - 1; i >= 0; i-- {
		result += c[i] * math.Pow(val, float64(i))
	}
	return result
}
