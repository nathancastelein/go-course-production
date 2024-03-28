package logging

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	user := c.QueryParam("user")
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", user))
}
