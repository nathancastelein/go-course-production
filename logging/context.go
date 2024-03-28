package logging

import (
	"context"
	"errors"
	"log/slog"
)

var (
	ErrLoggerNotFound    = errors.New("logger not found in context")
	ErrRequestIDNotFound = errors.New("request id not found in context")
)

func NewContextWithRequestID(ctx context.Context, requestId string) context.Context {
	return ctx
}

func GetRequestIDFromContext(ctx context.Context) (string, error) {
	return "", ErrRequestIDNotFound
}

func NewContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return ctx
}

func GetLoggerFromContext(ctx context.Context) (*slog.Logger, error) {
	return nil, ErrLoggerNotFound
}
