package numericalgo

import (
	"sort"
)

// CoordinatePair is a struct type with two properties, X and Y coordinates
type CoordinatePair struct {
	X float64
	Y float64
}

// SortCoordinatePairs is a function which receives a slice of CoordinatePairs, and sorts it in ascending order.
func SortCoordinatePairs(cp []CoordinatePair) {
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].X < cp[j].X
	})
}

// SlicesToCoordinatePairs is a function which receives two slices of floats (x and y), turns them into a slice of CoordinatePairs, and returns the result.
func SlicesToCoordinatePairs(x, y []float64) []CoordinatePair {
	cp := make([]CoordinatePair, len(x))
	for i := 0; i < len(x); i++ {
		cp[i] = CoordinatePair{X: x[i], Y: y[i]}
	}
	return cp
}
