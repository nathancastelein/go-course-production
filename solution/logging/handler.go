package logging

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	user := c.QueryParam("user")
	logger, err := GetLoggerFromContext(c.Request().Context())
	if err != nil {
		return err
	}

	logger.With(slog.String("user", user)).Info("saying hello")

	if user == "" {
		logger.Warn("empty user")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", user))
}
