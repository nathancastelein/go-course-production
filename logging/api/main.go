package main

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-production/logging"
)

func main() {
	logging.InitLogger()

	router := echo.New()
	router.Use(logging.Logger())
	router.GET("/", logging.Hello)

	if err := router.Start(":8080"); err != nil {
		slog.Error("router failed to start: %s", err)
	}
}
