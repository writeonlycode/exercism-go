package darts

import "math"

func Score(x, y float64) int {
	distanceFromCenter := math.Sqrt(math.Pow(math.Abs(x), 2) + math.Pow(math.Abs(y), 2))

	switch {
	case distanceFromCenter > 10:
		return 0
	case distanceFromCenter > 5:
		return 1
	case distanceFromCenter > 1:
		return 5
	case distanceFromCenter >= 0:
		return 10
	default:
		return -1
	}
}
