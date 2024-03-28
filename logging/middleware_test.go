package logging

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type log struct {
	Message     string `json:"msg"`
	HttpMethod  string `json:"http_method"`
	HttpURI     string `json:"http_uri"`
	RequestID   string `json:"request_id"`
	ElapsedTime int64  `json:"elapsed_time"`
}

type LoggerMiddlewareSuite struct {
	suite.Suite

	ctx        echo.Context
	buffer     bytes.Buffer
	middleware echo.HandlerFunc
	request    *http.Request
}

func (l *LoggerMiddlewareSuite) SetupSuite() {
	// Intercept logs
	logger := slog.New(slog.NewJSONHandler(&l.buffer, nil))
	slog.SetDefault(logger)

	// Prepare middleware
	l.middleware = Logger()(func(c echo.Context) error {
		return nil
	})
}

func (l *LoggerMiddlewareSuite) SetupTest() {
	// Prepare request
	l.request = httptest.NewRequest(http.MethodGet, "/", nil)
	l.request.Header.Add("X-Request-Id", uuid.New().String())

	// Prepare response recorder
	rec := httptest.NewRecorder()

	// Prepare context
	e := echo.New()
	l.ctx = e.NewContext(l.request, rec)

	// Reset buffer
	l.buffer.Reset()
}

func (l *LoggerMiddlewareSuite) TestSimpleLog() {
	// Arrange

	// Act
	err := l.middleware(l.ctx)

	// Assert
	l.Require().NoError(err)
	for _, expectedLog := range []string{
		"handle request",
		"request handled",
	} {
		receivedLog := &log{}
		str, err := l.buffer.ReadBytes('\n')
		l.Require().NoError(err, "expected to read log \"%s\", go nothing", expectedLog)
		l.Require().NoError(json.Unmarshal(str, receivedLog), "expected log to be JSON formatted")
		l.Require().Equal(expectedLog, receivedLog.Message)
	}
}

func (l *LoggerMiddlewareSuite) TestLogWithAttributes() {
	// Arrange

	// Act
	err := l.middleware(l.ctx)

	// Assert
	l.Require().NoError(err)
	for _, expectedLog := range []string{
		"handle request",
		"request handled",
	} {
		receivedLog := &log{}
		str, err := l.buffer.ReadBytes('\n')
		l.Require().NoError(err, "expected to read log \"%s\", go nothing", expectedLog)
		l.Require().NoError(json.Unmarshal(str, receivedLog), "expected log to be JSON formatted")
		l.Require().Equal(expectedLog, receivedLog.Message)
		l.Require().Equal(http.MethodGet, receivedLog.HttpMethod)
		l.Require().Equal("/", receivedLog.HttpURI)
	}
}

func (l *LoggerMiddlewareSuite) TestLogWithRequestID() {
	// Arrange

	// Act
	err := l.middleware(l.ctx)

	// Assert
	l.Require().NoError(err)
	for _, expectedLog := range []string{
		"handle request",
		"request handled",
	} {
		receivedLog := &log{}
		str, err := l.buffer.ReadBytes('\n')
		l.Require().NoError(err, "expected to read log \"%s\", go nothing", expectedLog)
		l.Require().NoError(json.Unmarshal(str, receivedLog), "expected log to be JSON formatted")
		l.Require().Equal(expectedLog, receivedLog.Message)
		l.Require().Equal(l.request.Header.Get("X-Request-Id"), receivedLog.RequestID)
	}
}

func (l *LoggerMiddlewareSuite) TestLogWithElapsedTime() {
	// Arrange
	middleware := Logger()(func(c echo.Context) error {
		time.Sleep(20 * time.Millisecond)
		return nil
	})

	// Act
	err := middleware(l.ctx)

	// Assert
	l.Require().NoError(err)
	for _, expectedLog := range []string{
		"handle request",
		"request handled",
	} {
		str, err := l.buffer.ReadBytes('\n')
		l.Require().NoError(err, "expected to read log \"%s\", go nothing", expectedLog)
		if expectedLog == "handle request" {
			// skip first log
			continue
		}
		receivedLog := &log{}
		l.Require().NoError(json.Unmarshal(str, receivedLog), "expected log to be JSON formatted")
		l.Require().Equal(expectedLog, receivedLog.Message)
		l.Require().NotZero(receivedLog.ElapsedTime, "elapsed time should not be equal to zero")
	}
}

func TestLoggerMiddlewareSuite(t *testing.T) {
	suite.Run(t, &LoggerMiddlewareSuite{})
}
