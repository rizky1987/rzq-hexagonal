package echo_middleware

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func TrackerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		trackerID := generateTrackerId()
		req := c.Request()
		// masukkan ke context.Context
		ctx := context.WithValue(req.Context(), "X-Tracker-Id", trackerID)
		c.SetRequest(req.WithContext(ctx))

		c.Response().Header().Set("X-Tracker-Id", trackerID)
		return next(c)
	}
}

func generateTrackerId() string {

	return fmt.Sprintf("%s-%s", os.Getenv("APP_ID"), time.Now().Format("20060102150405000"))
}
