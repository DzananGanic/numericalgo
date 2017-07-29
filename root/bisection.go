package root

import (
	"math"
)

// Bisection receives three parameters. First parameter is the function we want to find root of. Second one is the tolerance. Third and fourth parameters are left and right bounds of the function. The Bisection function returns the result as a float and error (if there is any).
func Bisection(f func(float64) float64, eps, l, r float64) (float64, error) {
	var result float64

	mid := (l + r) / 2

	if math.Abs(f(mid)) < eps {
		result = mid
	} else if (f(l) < 0) == (f(mid) < 0) {
		result, err := Bisection(f, eps, mid, r)
		if err != nil {
			return result, err
		}
	} else if (f(r) < 0) == (f(mid) < 0) {
		result, err := Bisection(f, eps, l, mid)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
