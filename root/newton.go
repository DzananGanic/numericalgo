package root

import "github.com/DzananGanic/numericalgo/differentiate"

// Newton receives three parameters. First parameter is the function we want to find root of. Second one is the initial guess (reasonably close to the true root). Third one is the number of iterations for the newton method. The Newton function returns the result as a float64, and the error.
func Newton(f func(float64) float64, x0 float64, iter int) (float64, error) {
	for i := 0; i < iter; i++ {
		d, err := differentiate.Central(f, x0, 0.01)
		if err != nil {
			return 0, err
		}
		x0 = x0 - f(x0)/d
	}

	return x0, nil
}
