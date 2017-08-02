package poly

import (
	"math"

	"github.com/DzananGanic/numericalgo"
)

// Poly type fits two vectors x and y, finds the appropriate coefficients and predicts the value such that y=p1+p2*x+p3*x^2+...+p(n+1)*x^n is the best approximation of the given data in a sense of the least square error.
type Poly struct {
	x     numericalgo.Vector
	y     numericalgo.Vector
	Coeff numericalgo.Vector
}

// New returns the pointer to the new Poly type
func New() *Poly {
	pf := &Poly{}
	return pf
}

// Fit function in Poly type receives two vectors, finds and stores the coefficients in the coeff property, and returns the error if something went wrong. Coefficients are calculated based on the y=p1+p2*x+p3*x^2+...+p(n+1)*x^n formula.
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

	p.Coeff, err = coeff.GetColumnAt(0)

	if err != nil {
		return err
	}

	return nil
}

// Predict function in Linear type accepts value to be predicted, and returns the predicted value based on the y=p1+p2*x+p3*x^2+...+p(n+1)*x^n formula.
func (p *Poly) Predict(val float64) float64 {
	var result float64
	c := p.Coeff
	for i := 0; i < len(c); i++ {
		result += c[i] * math.Pow(val, float64(i))
	}
	return result
}
