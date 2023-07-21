// Take any positive integer n. If n is even, divide n by 2. If n is odd,
// multiply n by 3 and add 1. Repeat the process indefinitely. Collatz
// Conjecture states that no matter which number you start with, you will
// always reach 1 eventually.
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1 from the
// given number n.
func CollatzConjecture(n int) (int, error) {
	return collatzConjectureHelper(n, 0)
}

func collatzConjectureHelper(n int, acc int) (int, error) {
	switch {
	case n < 1:
		return 0, errors.New("n must be >= 1")

	case n == 1:
		return acc, nil

	case n%2 == 0:
		return collatzConjectureHelper(n/2, acc+1)

	default:
		return collatzConjectureHelper((3*n)+1, acc+1)
	}
}
