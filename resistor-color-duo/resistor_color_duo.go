// Package resistorcolorduo defines functions to decode the resistance value of
// resistors.
package resistorcolorduo

import (
	"math"
)

var resistance = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value returns the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	input := colors[:2]
	result := 0

	for i := 0; i < len(input); i++ {
		result += resistance[input[i]] * (int(math.Pow(10, float64(len(input)-1-i))))
	}

	return result
}
