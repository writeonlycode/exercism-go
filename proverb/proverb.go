// Package proverb provides a function to generate a proverb from a list of
// words.
package proverb

import "fmt"

// Proverb returns a proverb generated from the given list of words.
func Proverb(rhyme []string) []string {
	var result []string

	if len(rhyme) == 0 {
		return result
	}

	for i := 1; i < len(rhyme); i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}

	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return result
}
