package numericalgo

import (
	"fmt"
	"math"
)

type Vector []float64

func (v Vector) Dim() int {
	return len(v)
}

func (v Vector) IsDimEqual(v2 Vector) bool {
	return v.Dim() == v2.Dim()
}

func (v Vector) IsSimilar(v2 Vector, tolerance float64) bool {
	if !v.IsDimEqual(v2) {
		return false
	}

	for i := range v {
		if math.Abs(v[i]-v2[i]) > tolerance {
			return false
		}
	}

	return true
}

func (v Vector) Power(n float64) Vector {
	var result Vector

	for _, val := range v {
		result = append(result, math.Pow(val, n))
	}

	return result
}

func (v Vector) Add(v2 Vector) (Vector, error) {
	var result Vector

	if !v.IsDimEqual(v2) {
		return result, fmt.Errorf("Dimensions must match")
	}

	for index := range v {
		result = append(result, v[index]+v2[index])
	}
	return result, nil
}

func (v Vector) Subtract(v2 Vector) (Vector, error) {
	var result Vector

	if !v.IsDimEqual(v2) {
		return result, fmt.Errorf("Dimensions must match")
	}

	for index := range v {
		result = append(result, v[index]-v2[index])
	}
	return result, nil
}

func (v Vector) Dot(v2 Vector) (float64, error) {
	var result float64

	if !v.IsDimEqual(v2) {
		return result, fmt.Errorf("Dimensions must match")
	}

	for index := range v {
		result += v[index] * v2[index]
	}

	return result, nil
}

func (v Vector) MultiplyByScalar(s float64) Vector {
	var result Vector

	for index := range v {
		result = append(result, v[index]*s)
	}
	return result
}

func (v Vector) DivideByScalar(s float64) (Vector, error) {
	var result Vector

	if s == 0 {
		return result, fmt.Errorf("Cannot divide by zero")
	}

	for index := range v {
		result = append(result, v[index]/s)
	}
	return result, nil
}
