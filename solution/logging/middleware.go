package logging

import (
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestId := c.Request().Header.Get("X-Request-Id")
			if requestId == "" {
				requestId = uuid.New().String()
			}

			logger := slog.With(
				slog.String("http_method", c.Request().Method),
				slog.String("http_uri", c.Request().RequestURI),
				slog.String("request_id", requestId),
			)

			ctx := c.Request().Context()
			ctx = NewContextWithRequestID(ctx, requestId)
			ctx = NewContextWithLogger(ctx, logger)
			c.SetRequest(c.Request().WithContext(ctx))

			logger.Info("handle request")
			start := time.Now()
			if err := next(c); err != nil {
				return err
			}

			logger.Info("request handled", slog.Int64("elapsed_time", time.Since(start).Milliseconds()))
			return nil
		}
	}
}

func InitLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).
		With(
			slog.String("application", "api"),
		)
	slog.SetDefault(logger)
}
