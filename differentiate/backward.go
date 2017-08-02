package differentiate

import (
	"fmt"
)

func Backward(f func(float64) float64, val, h float64) (float64, error) {
	if h <= 0 {
		return 0, fmt.Errorf("Step size has to be greater than 0")
	}
	return (f(val) - f(val-h)) / h, nil
}
