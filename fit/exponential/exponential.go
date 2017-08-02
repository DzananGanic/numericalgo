package exponential

import (
	"github.com/DzananGanic/numericalgo"
)

// Exponential type fits two vectors x and y, finds the appropriate coefficients and predicts the value such that y=p*e^(q*x) is the best approximation of the given data in a sense of the least square error.
type Exponential struct {
	x     numericalgo.Vector
	y     numericalgo.Vector
	Coeff numericalgo.Vector
}

// New returns the pointer to the new Exponential type
func New() *Exponential {
	ef := &Exponential{}
	return ef
}

// Fit function in Exponential type receives two vectors, finds and stores the coefficients in the coeff property, and returns the error if something went wrong. Coefficients are calculated based on the y=p*e^(q*x) formula.
func (e *Exponential) Fit(x numericalgo.Vector, y numericalgo.Vector) error {
	xMatrix := numericalgo.Matrix{x}
	yMatrix := numericalgo.Matrix{y}

	xLogMatrix := xMatrix.Log()
	yLogMatrix := yMatrix.Log()

	xLogT, err := xLogMatrix.Transpose()
	if err != nil {
		return err
	}

	yLogT, err := yLogMatrix.Transpose()
	if err != nil {
		return err
	}

	ones := make(numericalgo.Vector, x.Dim())
	for i := range ones {
		ones[i] = 1
	}

	X, err := xLogT.AddColumnAt(0, ones)

	if err != nil {
		return err
	}

	coeff, err := X.LeftDivide(yLogT)

	if err != nil {
		return err
	}

	expCoeff := coeff.Exp()
	e.Coeff, err = expCoeff.GetColumnAt(0)

	if err != nil {
		return err
	}

	return nil
}

// Predict function in Exponential type accepts value to be predicted, and returns the predicted value based on the y=p*e^(q*x) formula.
func (e *Exponential) Predict(val float64) float64 {
	c := e.Coeff
	return c[1]*val + c[0]
}
