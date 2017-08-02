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
	var r []float64
	for _, val := range vals {
		est, err := WithSingle(vi, val)
		if err != nil {
			return r, err
		}
		r = append(r, est)
	}
	return r, nil
}

// WithSingle accepts the single float64 value, and returns the interpolated value for it, and the error
func WithSingle(vi validateInterpolator, val float64) (float64, error) {
	var est float64

	err := vi.Validate(val)
	if err != nil {
		return est, err
	}

	est = vi.Interpolate(val)
	return est, nil
}
