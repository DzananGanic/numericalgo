package integrate

import (
	"fmt"

	"github.com/DzananGanic/numericalgo"
)

// Simpson is a function which accepts function, left, right bounds and n number of subdivisions. It returns the integration
// value of the function in the given bounds using simpson rule.
func Simpson(f func(float64) float64, l, r float64, n int) (float64, error) {

	var eval, evalOdd, evalEven numericalgo.Vector
	var x float64

	if n == 0 {
		return 0, fmt.Errorf("Number of subdivisions n cannot be 0")
	}

	h := (r - l) / float64(n)

	for i := 0; i <= n; i++ {
		x = l + h*float64(i)
		eval = append(eval, f(x))
	}

	for i := 1; i < n; i++ {
		if i%2 == 0 {
			evalEven = append(evalEven, eval[i])
		} else {
			evalOdd = append(evalOdd, eval[i])
		}
	}

	return h / 3 * (eval[0] + eval[n] + 4*evalOdd.Sum() + 2*evalEven.Sum()), nil
}
