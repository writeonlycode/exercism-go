// Package luhn provides a function to determine whether a given id is a valid
// according to Luhn's formula.
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns true if the given id is a valid according to Luhn's formula.
func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	sum := 0

	for i := 1; i <= len(id); i++ {
		index := len(id) - i

		value, err := strconv.Atoi(string(id[index]))

		if err != nil {
			return false
		}

		if i > 1 && i%2 == 0 {
			value *= 2
		}

		if value > 9 {
			value -= 9
		}

		sum += value
	}

	return sum%10 == 0
}
