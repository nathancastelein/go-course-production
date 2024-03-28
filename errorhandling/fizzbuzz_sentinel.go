package errorhandling

import "fmt"

func FizzbuzzWithSentinelError(input int) (string, error) {
	if input == 0 {
		return "", fmt.Errorf("input cannot be equal to 0: %d", input)
	}

	if input < 0 {
		return "", fmt.Errorf("input cannot be negative: %d", input)
	}

	return Fizzbuzz(input), nil
}

func MyFizzbuzzProcess(input int) error {
	// Call FizzbuzzWithSentinelError
	// If error is a zero input error, return nil
	// If error is a negative input error, return the error
	// Otherwise print the result
	return nil
}
