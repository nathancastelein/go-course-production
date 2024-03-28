package errorhandling

import (
	"errors"
	"fmt"

	"github.com/nathancastelein/go-course-production/errorhandling"
)

type InvalidInput struct {
	input int
}

func NewInvalidInput(input int) error {
	return InvalidInput{
		input: input,
	}
}

func (i InvalidInput) Error() string {
	return fmt.Sprintf("got invalid input %d", i.input)
}

func FizzbuzzWithCustomError(input int) (string, error) {
	if input <= 0 {
		return "", NewInvalidInput(input)
	}

	return errorhandling.Fizzbuzz(input), nil
}

func MyFizzbuzzProcessWithCustomError(input int) error {
	result, err := FizzbuzzWithCustomError(input)
	var errInvalidInput InvalidInput
	if err != nil {
		if errors.As(err, &errInvalidInput) {
			return nil
		} else {
			return err
		}
	}

	fmt.Printf("Fizzbuzz of %d is %s\n", input, result)
	return nil
}
