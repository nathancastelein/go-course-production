package errorhandling

import (
	"errors"
	"fmt"

	"github.com/nathancastelein/go-course-production/errorhandling"
)

var (
	ErrZeroInput     = errors.New("invalid input, can not be zero")
	ErrNegativeInput = errors.New("invalid input, can not be negative")
)

func FizzbuzzWithSentinelError(input int) (string, error) {
	if input == 0 {
		return "", ErrZeroInput
	}

	if input < 0 {
		return "", ErrNegativeInput
	}

	return errorhandling.Fizzbuzz(input), nil
}

func MyFizzbuzzProcess(input int) error {
	result, err := FizzbuzzWithSentinelError(input)
	if err != nil {
		if errors.Is(err, ErrZeroInput) {
			return nil
		} else if errors.Is(err, ErrNegativeInput) {
			return err
		} else {
			return err
		}
	}

	fmt.Printf("Fizzbuzz of %d is %s\n", input, result)
	return nil
}
