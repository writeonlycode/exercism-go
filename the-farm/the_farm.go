package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood calculates the amount of fodder per cow.
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	fa, faErr := fc.FodderAmount(cows)

	if faErr != nil {
		return 0, faErr
	}

	ff, ffErr := fc.FatteningFactor()

	if ffErr != nil {
		return 0, ffErr
	}

	return (fa / float64(cows)) * ff, nil
}

// ValidateInputAndDivideFood validates the number of cows given as input, and
// calculates the amount of fodder per cow.
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, cows)
}

// InvalidCowsError represents an error when the number of cows is invalid.
type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// ValidateNumberOfCows validates the number of cows given as input, and
// returns an error it the input is invalid.
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	}

	return nil
}
