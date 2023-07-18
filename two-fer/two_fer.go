// Package twofer implements implements functions to determine what you will
// say as you give away a cookie to someone.
package twofer

import "fmt"

// ShareWith takes a name and returns the string "One for {name}, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
