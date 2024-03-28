package logging

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestIDInContext(t *testing.T) {
	// Arrange
	expectedRequestId := "foo"

	// Act
	requestId, err := GetRequestIDFromContext(NewContextWithRequestID(context.Background(), expectedRequestId))

	// Assert
	require.NoError(t, err)
	require.Equal(t, expectedRequestId, requestId)
}

func TestLoggerInContext(t *testing.T) {
	// Arrange
	expectedLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Act
	requestId, err := GetLoggerFromContext(NewContextWithLogger(context.Background(), expectedLogger))

	// Assert
	require.NoError(t, err)
	require.Equal(t, expectedLogger, requestId)
}
