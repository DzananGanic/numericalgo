package root

import (
	"math"
)

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
