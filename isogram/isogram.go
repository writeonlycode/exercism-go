// Package isogram is a small package that provides a function to determine
// whether a word is an isogram.
package isogram

import "unicode"

// IsIsogram returns true if the given word is an isogram.
func IsIsogram(word string) bool {
	letters := map[rune]bool{}

	for _, v := range word {
		if letters[unicode.ToUpper(v)] {
			return false
		}

		if v != ' ' && v != '-' {
			letters[unicode.ToUpper(v)] = true
		}
	}

	return true
}
