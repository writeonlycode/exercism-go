// Package diffsquares provide functions to calculate the square of sums, the
// sum of squares, and the difference between these of the first n natural
// numbers.
package diffsquares

// SquareOfSum returns the square of the sum of the first n natural nuebers.
func SquareOfSum(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares returns the sum of the square of the first n natural numbers.
func SumOfSquares(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i * i
	}

	return sum

}

// Difference returns the difference between the square of the sum and the sum
// of the squares of the firs n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
