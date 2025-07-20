package echo_middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	middlewareHelper "rzq-hexagonal/cmd/http/middleware/helper"
	"time"

	"github.com/labstack/echo/v4"
)

func GetCallerInfo() string {
	for i := 2; i < 10; i++ { // skip internal runtime frames
		pc, file, line, ok := runtime.Caller(i)
		if ok && !isInternal(file) {
			fn := runtime.FuncForPC(pc)
			return fmt.Sprintf("%s:%d (%s)", file, line, fn.Name())
		}
	}
	return "unknown"
}

func isInternal(file string) bool {
	return filepath.Base(file) == "echo.go" || filepath.Base(file) == "fiber.go"
}

func ErrorRequestLogging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			start := time.Now()

			err := next(c)
			status := c.Response().Status
			if status != http.StatusOK {
				// Headers
				headers := map[string][]string{}
				for k, v := range c.Request().Header {
					headers[k] = v
				}

				// Body
				var requestBody []byte
				if c.Request().Body != nil {
					bodyBytes, err := io.ReadAll(c.Request().Body)
					if err == nil {
						requestBody = bodyBytes
						// Step 2: Replace body agar bisa dibaca ulang oleh handler
						c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
					}
				}

				fmt.Println("rizky")
				fmt.Println(string(requestBody))
				// Capture response
				rc := &middlewareHelper.ResponseCaptureWriter{
					ResponseWriter: c.Response().Writer,
					Body:           &bytes.Buffer{},
				}
				c.Response().Writer = rc

				maskedResponseBody := middlewareHelper.MaskSensitiveData(rc.Body.Bytes())
				traceID := c.Get("TraceID")
				callerInfo := GetCallerInfo()
				logMsg := "[Echo] %s %s [%d] (%v) TraceID=%v\nHeaders=%v\nBody=%s\nCaller: %s\nResponse Body %s\n"
				log.Printf(logMsg,
					c.Request().Method,
					c.Request().URL.Path,
					status,
					time.Since(start),
					traceID,
					headers,
					"maskedReqBody",
					callerInfo,
					maskedResponseBody,
				)
			}
			return err
		}
	}
}
