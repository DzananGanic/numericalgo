package fit

type predictor interface {
	Predict(float64) float64
}

// PredictMulti accepts the slice of float64, and returns the predicted values for the passed slice values, and the error
func PredictMulti(p predictor, vals []float64) []float64 {
	var r []float64
	for _, val := range vals {
		est := p.Predict(val)
		r = append(r, est)
	}
	return r
}
