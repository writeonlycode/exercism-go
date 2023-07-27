// Package pangram defines a function to check whethera sentence is a  pangram.
package pangram

import (
	"strings"
)

// IsPangram returns true if the sentence is a pangram.
func IsPangram(input string) bool {
	sentence := strings.ToLower(input)

	for letter := 'a'; letter <= 'z'; letter++ {
		if !strings.ContainsRune(sentence, letter) {
			return false
		}
	}

	return true
}
