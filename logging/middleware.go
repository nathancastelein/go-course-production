package logging

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				return err
			}

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
