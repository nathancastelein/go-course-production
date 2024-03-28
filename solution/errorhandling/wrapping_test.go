package errorhandling

import (
	"errors"
	"testing"
)

func TestFizzbuzzWithWrappingError(t *testing.T) {
	// Arrange
	for _, input := range []int{0, -3} {
		// Act
		_, err := FizzbuzzWithWrappingError(input)

		// Assert
		if err == nil {
			t.Fatal("test failed, expecting error but got nil")
		}

		if !errors.Is(err, ErrInvalidInput) {
			t.Fatalf("test failed, expecting error type ErrInvalidInput")
		}
	}
}
