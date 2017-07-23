package numericalgo

import (
	"sort"
)

// CoordinatePair is a struct type with two properties, X and Y coordinates
type CoordinatePair struct {
	X float64
	Y float64
}

func SortCoordinatePairs(pairs []CoordinatePair) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].X < pairs[j].X
	})
}

func SlicesToCoordinatePairs(x, y []float64) []CoordinatePair {
	coordinatePairs := make([]CoordinatePair, len(x))
	for i := 0; i < len(x); i++ {
		coordinatePairs = append(coordinatePairs, CoordinatePair{X: x[i], Y: y[i]})
	}
	return coordinatePairs
}
