package errorhandling

import (
	"fmt"

	"github.com/nathancastelein/go-course-production/errorhandling"
)

func FizzbuzzWithError(input int) (string, error) {
	if input == 0 {
		return "", fmt.Errorf("input cannot be equal to 0: %d", input)
	}

	if input < 0 {
		return "", fmt.Errorf("input cannot be negative: %d", input)
	}

	return errorhandling.Fizzbuzz(input), nil
}
