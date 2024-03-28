package logging

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmailLogValuer(t *testing.T) {
	var buffer bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buffer, nil))
	email := Email("john.doe@example.com")
	type log struct {
		Email string `json:"email"`
	}

	logger.Info("hello", slog.Any("email", email))

	//
	logLine := buffer.String()
	receivedLog := &log{}
	require.NoError(t, json.Unmarshal([]byte(logLine), receivedLog), "expected log to be JSON formatted")
	require.Equal(t, "*******@******.***", receivedLog.Email)
}
