package errorhandling

import (
	"reflect"
	"testing"
)

func TestFizzBuzzWithCustomError(t *testing.T) {
	// Arrange
	for _, input := range []int{0, -3} {
		// Act
		_, err := FizzbuzzWithCustomError(input)

		// Assert
		if err == nil {
			t.Fatal("test failed, expecting error but got nil")
		}

		if reflect.TypeOf(err).Name() != "InvalidInput" {
			t.Fatalf("test failed, expecting error type InvalidInput, got %s", reflect.TypeOf(err).Name())
		}
	}
}
