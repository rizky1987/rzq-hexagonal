package echo_middleware

import (
	"github.com/labstack/echo/v4"
	echoLimiter "github.com/labstack/echo/v4/middleware"
)

func EchoRateLimiter() echo.MiddlewareFunc {
	return echoLimiter.RateLimiterWithConfig(echoLimiter.RateLimiterConfig{
		Skipper: echoLimiter.DefaultSkipper,
		Store:   echoLimiter.NewRateLimiterMemoryStore(20), // 20 req / sec per IP (default)
	})
}
