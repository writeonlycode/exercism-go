// Package reverse defines a function to reverse a string.
package reverse

// Reverse takes a string and returns the reversed string.
func Reverse(input string) string {
	reversed := ""

	for _, v := range input {
		reversed = string(v) + reversed
	}

	return reversed
}
