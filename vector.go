package numericalgo

import (
	"fmt"
	"math"
)

// Vector is the []float64 which has custom methods needed for vector operations.
type Vector []float64

// Dim returns the dimension of the vector.
func (v Vector) Dim() int {
	return len(v)
}

// AreDimsEqual receives another vector as a parameter. Method returns true if the dimensions of the vectors are equal.
func (v Vector) AreDimsEqual(v2 Vector) bool {
	return v.Dim() == v2.Dim()
}

// IsSimilar receives another vector and tolerance as a parameter. It checks whether the two vectors are similar within the provided tolerance.
func (v Vector) IsSimilar(v2 Vector, tol float64) bool {

	if !v.AreDimsEqual(v2) {
		return false
	}

	for i := range v {
		if math.Abs(v[i]-v2[i]) > tol {
			return false
		}
	}

	return true
}

// Sum returns the sum of all elements in the vector
func (v Vector) Sum() float64 {
	var sum float64
	for _, val := range v {
		sum += val
	}
	return sum
}

// Power receives a float as a parameter. It returns the vector whose elements are x^n.
func (v Vector) Power(n float64) Vector {

	for i, val := range v {
		v[i] = math.Pow(val, n)
	}

	return v
}

// Add receives another vector as a parameter. It adds the two vectors and returns the result vector and the error (if there is any).
func (v Vector) Add(v2 Vector) (Vector, error) {
	var r Vector

	if !v.AreDimsEqual(v2) {
		return r, fmt.Errorf("Dimensions must match")
	}

	for index := range v {
		r = append(r, v[index]+v2[index])
	}

	return r, nil
}

// Subtract receives another vector as a parameter. It subtracts the two vectors and returns the result vector and an error (if there is any).
func (v Vector) Subtract(v2 Vector) (Vector, error) {
	var r Vector

	if !v.AreDimsEqual(v2) {
		return r, fmt.Errorf("Dimensions must match")
	}

	for index := range v {
		r = append(r, v[index]-v2[index])
	}

	return r, nil
}

// Dot receives another vector as a parameter. It calculates the dot product between the two vectors and returns the float result and an error (if there is any).
func (v Vector) Dot(v2 Vector) (float64, error) {
	var r float64

	if !v.AreDimsEqual(v2) {
		return r, fmt.Errorf("Dimensions must match")
	}

	for index := range v {
		r += v[index] * v2[index]
	}

	return r, nil
}

// MultiplyByScalar receives a scalar as a parameter. It multiplies all the elements of the vector with provided scalar and returns the result vector.
func (v Vector) MultiplyByScalar(s float64) Vector {
	var r Vector

	for index := range v {
		r = append(r, v[index]*s)
	}

	return r
}

// DivideByScalar receives a scalar as a parameter. It divides all the elements of the vector by provided scalar and returns the result vector.
func (v Vector) DivideByScalar(s float64) (Vector, error) {
	var r Vector

	if s == 0 {
		return r, fmt.Errorf("Cannot divide by zero")
	}

	for index := range v {
		r = append(r, v[index]/s)
	}

	return r, nil
}
