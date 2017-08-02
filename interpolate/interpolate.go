package interpolate

type interpolator interface {
	Interpolate(float64) float64
}

type validator interface {
	Validate(float64) error
}

type validateInterpolator interface {
	interpolator
	validator
}

// WithMulti accepts the slice of float64, and returns the interpolated values for the passed slice values, and the error
func WithMulti(vi validateInterpolator, vals []float64) ([]float64, error) {
	var result []float64
	for _, val := range vals {
		estimate, err := WithSingle(vi, val)
		if err != nil {
			return result, err
		}
		result = append(result, estimate)
	}
	return result, nil
}

// WithSingle accepts the single float64 value, and returns the interpolated value for it, and the error
func WithSingle(vi validateInterpolator, val float64) (float64, error) {
	var estimate float64

	err := vi.Validate(val)
	if err != nil {
		return estimate, err
	}

	estimate = vi.Interpolate(val)
	return estimate, nil
}
