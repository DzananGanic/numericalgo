package integrate

import (
	"fmt"

	"github.com/teivah/numericalgo"
)

// Trapezoid is a function which accepts function, left, right bounds and n number of subdivisions. It returns the integration
// value of the function in the given bounds using trapezoidal rule.
func Trapezoid(f func(float64) float64, l, r float64, n int) (float64, error) {

	var eval numericalgo.Vector
	var x float64

	if n == 0 {
		return 0, fmt.Errorf("Number of subdivisions n cannot be 0")
	}

	h := (r - l) / float64(n)

	for i := 0; i <= n; i++ {
		x = l + h*float64(i)
		eval = append(eval, f(x))
	}

	return h * ((eval[0]+eval[n])/2 + eval[1:n].Sum()), nil
}
