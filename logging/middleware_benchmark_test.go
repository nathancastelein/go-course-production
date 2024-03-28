package logging

import (
	"io/ioutil"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func BenchmarkLogging(b *testing.B) {
	// Arrange
	// Drop logs
	logger := slog.New(slog.NewJSONHandler(ioutil.Discard, nil))
	slog.SetDefault(logger)

	// Prepare a new request request
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	// Prepare response recorder
	rec := httptest.NewRecorder()
	// Prepare context
	ctx := echo.New().NewContext(request, rec)
	// Prepare middleware
	middleware := Logger()(func(c echo.Context) error {
		return nil
	})

	// Act
	for n := 0; n < b.N; n++ {
		middleware(ctx)
	}
}
