// Pacakge etl defines a function to convert betweek two ways of store point
// values of different letters.
package etl

import "strings"

// Transform takes a map from points to a list of letters, and converts it to a
// map from letters to their point values.
func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}

	for points, letters := range in {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = points
		}
	}

	return out
}
