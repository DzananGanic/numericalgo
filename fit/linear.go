package fit

import (
	"github.com/DzananGanic/numericalgo"
)

// Linear is a type which embeds Base type. Linear type fits two vectors x and y, finds the appropriate coefficients and predicts the value such that y=p+qx is the best approximation of the given data in a sense of the least square error.
type Linear struct {
	Base
}

// NewLinear returns the pointer to the new Linear type
func NewLinear() *Linear {
	lf := &Linear{}
	return lf
}

// Fit function in Linear type receives two vectors, finds and stores the coefficients in the coeff property, and returns the error if something went wrong. Coefficients are calculated based on the y=p+q*x formula.
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

	coeff, err := X.LeftDivide(Y)

	if err != nil {
		return err
	}

	l.coeff, err = coeff.GetColumnAt(0)

	if err != nil {
		return err
	}

	return nil
}

// Predict function in Linear type accepts value to be predicted, and returns the predicted value based on the y=p+q*x formula.
func (l *Linear) Predict(val float64) float64 {
	c := l.Coeff()
	return c[0] + c[1]*val
}
