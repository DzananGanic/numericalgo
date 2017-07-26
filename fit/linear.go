package fit

import (
	"github.com/DzananGanic/numericalgo"
)

type Linear struct {
	Base
}

func NewLinear() *Linear {
	lf := &Linear{}
	return lf
}

func (l *Linear) Fit(x numericalgo.Vector, y numericalgo.Vector) error {
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

	Y, err := yMatrix.Transpose()

	if err != nil {
		return err
	}

	coeffs, err := X.LeftDivide(Y)

	if err != nil {
		return err
	}

	l.q = coeffs[0][0]
	l.p = coeffs[1][0]

	return nil
}
