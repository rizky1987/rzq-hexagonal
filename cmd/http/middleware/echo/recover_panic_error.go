package echo_middleware

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

func EchoRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				if rec := recover(); rec != nil {
					traceID := c.Get("TraceID")
					method := c.Request().Method
					path := c.Request().URL.Path
					stack := string(debug.Stack())

					msg := fmt.Sprintf("panic: %v", rec)

					log.Printf("[PANIC][Echo] %s %s TraceID=%v\n%s msg %s", method, path, traceID, stack, msg)

					c.Error(echo.NewHTTPError(500, "Internal Server Error"))
				}
			}()

			return next(c)
		}
	}
}
