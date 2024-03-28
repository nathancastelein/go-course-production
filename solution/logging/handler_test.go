package logging

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func (l *LoggerMiddlewareSuite) TestLoggerFromContext() {
	// Arrange
	// Prepare a new request request
	l.request = httptest.NewRequest(http.MethodGet, "/", nil)
	l.request.Header.Add("X-Request-Id", "ac4c39f2-99c5-435a-ac4d-0fbacf85cbbd")
	query := l.request.URL.Query()
	query.Add("user", "Nathan")
	l.request.URL.RawQuery = query.Encode()

	// Prepare response recorder
	rec := httptest.NewRecorder()

	// Prepare context
	l.ctx = echo.New().NewContext(l.request, rec)

	middleware := Logger()(Hello)
	type logWithName struct {
		log
		Name string `json:"user"`
	}

	// Act
	err := middleware(l.ctx)

	// Assert
	l.Require().NoError(err)
	for _, expectedLog := range []string{
		"handle request",
		"saying hello",
		"request handled",
	} {
		str, err := l.buffer.ReadBytes('\n')
		l.Require().NoError(err, "expected to read log \"%s\", go nothing", expectedLog)
		if expectedLog != "saying hello" {
			// skip first log
			continue
		}
		receivedLog := &logWithName{}
		l.Require().NoError(json.Unmarshal(str, receivedLog), "expected log to be JSON formatted")
		l.Require().Equal(expectedLog, receivedLog.Message)
		l.Require().Equal("Nathan", receivedLog.Name)
	}
}

func (l *LoggerMiddlewareSuite) TestWarningWithNoQueryParam() {
	// Arrange
	middleware := Logger()(Hello)
	type logWithLevel struct {
		log
		Level string `json:"level"`
	}

	// Act
	err := middleware(l.ctx)

	// Assert
	l.Require().NoError(err)
	for _, expectedLog := range []string{
		"handle request",
		"saying hello",
		"empty user",
		"request handled",
	} {
		str, err := l.buffer.ReadBytes('\n')
		l.Require().NoError(err, "expected to read log \"%s\", go nothing", expectedLog)
		if expectedLog != "empty user" {
			continue
		}
		receivedLog := &logWithLevel{}
		l.Require().NoError(json.Unmarshal(str, receivedLog), "expected log to be JSON formatted")
		l.Require().Equal(expectedLog, receivedLog.Message)
		l.Require().Equal("WARN", receivedLog.Level)
	}
}
