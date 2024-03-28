package logging

import (
	"context"
	"errors"
	"log/slog"
)

type key int

const (
	requestIdKey key = 0
	loggerKey    key = 1
)

var (
	ErrLoggerNotFound    = errors.New("logger not found in context")
	ErrRequestIDNotFound = errors.New("request id not found in context")
)

func NewContextWithRequestID(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId)
}

func GetRequestIDFromContext(ctx context.Context) (string, error) {
	value := ctx.Value(requestIdKey)
	if value == nil {
		return "", ErrRequestIDNotFound
	}

	requestId, ok := value.(string)
	if !ok {
		return "", ErrRequestIDNotFound
	}
	return requestId, nil
}

func NewContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLoggerFromContext(ctx context.Context) (*slog.Logger, error) {
	value := ctx.Value(loggerKey)
	if value == nil {
		return nil, ErrLoggerNotFound
	}
	logger, ok := value.(*slog.Logger)
	if !ok {
		return nil, ErrLoggerNotFound
	}
	return logger, nil
}
