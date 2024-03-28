package errorhandling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyFizzbuzzProcessZero(t *testing.T) {
	// Arrange

	// Act
	err := MyFizzbuzzProcess(0)

	// Assert
	require.NoError(t, err)
}

func TestMyFizzbuzzProcessNegative(t *testing.T) {
	// Arrange

	// Act
	err := MyFizzbuzzProcess(-1)

	// Assert
	require.Error(t, err)
}
