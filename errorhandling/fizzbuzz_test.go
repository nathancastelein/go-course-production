package errorhandling

import (
	"strconv"
	"strings"
	"testing"
)

func TestFizzBuzzWithError(t *testing.T) {
	// Arrange
	type test struct {
		input            int
		isExpectingError bool
	}

	tests := map[string]test{
		"should return an error when number is 0": {
			input:            0,
			isExpectingError: true,
		},
		"should return an error when number is negative": {
			input:            -3,
			isExpectingError: true,
		},
	}

	// Act
	for name, test := range tests {
		_, err := FizzbuzzWithError(test.input)

		// Assert
		if test.isExpectingError && err == nil {
			t.Fatalf("%s failed, expecting error but got nil", name)
		}

		if !test.isExpectingError && err != nil {
			t.Fatalf("%s failed, not expecting error but got %s", name, err)
		}

		if err != nil {
			if !strings.Contains(err.Error(), strconv.Itoa(test.input)) {
				t.Fatalf("%s failed, expected to find %d in error message, got %s", name, test.input, err)
			}
		}
	}
}
