// Package grains defines functions to calculate the number of grains in a
// chessboard that has 64 squares, where square 1 has one grain, square 2 has
// two grains, and so on.
package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains in the given square.
func Square(number int) (uint64, error) {
	if number < 1 || 64 < number {
		return 0, errors.New("Number must be between 1 and 64.")
	}
	result := math.Pow(2, float64(number-1))
	return uint64(result), nil
}

// Total returns the total number of grains in the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
