package root

import (
	"math"
)

// Bisection receives three parameters. First parameter is the function we want to find root of. Second one is the tolerance. Third and fourth parameters are left and right bounds of the function.
func Bisection(f func(float64) float64, eps, l, r float64) float64 {

	mid := (l + r) / 2

	if math.Abs(f(mid)) < eps {
		return mid
	} else if (f(l) < 0) == (f(mid) < 0) {
		return Bisection(f, eps, mid, r)
	} else if (f(r) < 0) == (f(mid) < 0) {
		return Bisection(f, eps, l, mid)
	}

	return math.Inf(1)
}
