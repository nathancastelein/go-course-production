package errorhandling

import (
	"fmt"
)

func FizzbuzzWithError(input int) (string, error) {
	return Fizzbuzz(input), nil
}

func Fizzbuzz(input int) string {
	fizzBuzzValue := ""
	if input%3 == 0 {
		fizzBuzzValue += "Fizz"
	}

	if input%5 == 0 {
		fizzBuzzValue += "Buzz"
	}

	if fizzBuzzValue != "" {
		return fizzBuzzValue
	}

	return fmt.Sprintf("%d", input)
}
