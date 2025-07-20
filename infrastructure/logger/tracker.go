// internal/log/logger.go
package logger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type AppLogger struct {
	Logger  *slog.Logger
	Ctx     context.Context
	Request *http.Request
}

func New() *AppLogger {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler = slog.NewTextHandler(os.Stdout, opts)

	base := slog.New(handler)
	slog.SetDefault(slog.New(handler))
	return &AppLogger{
		Logger: base,
	}

}

func (al *AppLogger) LoggingBeginRequest(request *http.Request) {

	al.Ctx = request.Context()
	al.Request = request

	customMsg := fmt.Sprintf("[%s] %s. %s Headers : %s, Body : %s File : %s",
		al.getTrackerId(),
		"[Begin Request]",
		al.getRequestInfo(),
		al.getHeaderRequest(),
		al.getBodyRequest(),
		al.getCallerInfo())

	al.Logger.InfoContext(al.Ctx, customMsg)
}
func (al *AppLogger) LoggingEndRequest(request *http.Request) {

	al.Ctx = request.Context()
	al.Request = request

	fmt.Printf("[%s] %s. %s File : %s",
		al.getTrackerId(),
		"[End Request]",
		al.getRequestInfo(),
		al.getCallerInfo())

}

func (al *AppLogger) getCallerInfo() string {
	_, file, line, _ := runtime.Caller(2)

	return fmt.Sprintf("%s:%d", file, line)
}

func (al *AppLogger) getRequestInfo() string {
	method := ""
	path := ""
	if al.Request != nil {
		method = al.Request.Method
		path = al.Request.URL.Path
	}

	return fmt.Sprintf("[%s] %s", method, path)
}

func (al *AppLogger) getTrackerId() string {
	trackerId, _ := al.Ctx.Value("X-Tracker-Id").(string)

	return trackerId
}

func (al *AppLogger) getHeaderRequest() map[string]string {
	// Ambil body (dengan clone supaya nggak rusak)
	headers := map[string]string{}
	for k, v := range al.Request.Header {
		headers[k] = strings.Join(v, ",")
	}

	return headers
}

func (al *AppLogger) getBodyRequest() string {
	// Ambil body (dengan clone supaya nggak rusak)
	var bodyStr string
	if al.Request.Body != nil {
		bodyBytes, _ := io.ReadAll(al.Request.Body)
		bodyStr = strings.ReplaceAll(strings.ReplaceAll(string(bodyBytes), "\n", " "), "\r", " ")
		al.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return bodyStr
}
