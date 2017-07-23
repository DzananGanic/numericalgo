package interpolation

// Interpolator interface has one unimplemented method - Interpolate, which receives valueToInterpolate float64 parameter, and returns the estimate and error
type Interpolator interface {
	Interpolate(valueToInterpolate float64) float64
}

// Validator interface has one unimplemented method - Validate, which receives valueToInterpolate, and checks whether there are input errors.
type Validator interface {
	validate(valueToInterpolate float64) error
}

// ValidatedInterpolator interface is the compound interface which merges Interpolator and Validator together
type ValidatedInterpolator interface {
	Interpolator
	Validator
}

// InterpolateMultipleValues accepts the slice of float64, and returns the interpolated values for the passed slice values, and the error
func InterpolateMultipleValues(validatedInterpolator ValidatedInterpolator, valuesToInterpolate []float64) ([]float64, error) {
	var result []float64
	for _, valueToInterpolate := range valuesToInterpolate {
		estimate, err := InterpolateSingleValue(validatedInterpolator, valueToInterpolate)
		if err != nil {
			return result, err
		}
		result = append(result, estimate)
	}
	return result, nil
}

// InterpolateSingleValue accepts the single float64 value, and returns the interpolated value for it, and the error
func InterpolateSingleValue(validatedInterpolator ValidatedInterpolator, valueToInterpolate float64) (float64, error) {
	var estimate float64

	err := validatedInterpolator.validate(valueToInterpolate)
	if err != nil {
		return estimate, err
	}

	estimate = validatedInterpolator.Interpolate(valueToInterpolate)
	return estimate, nil
}
