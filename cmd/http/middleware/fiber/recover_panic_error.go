package fiber_middleware

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

func FiberRecover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if rec := recover(); rec != nil {
				traceID := c.Locals("TraceID")
				method := c.Method()
				path := c.Path()
				stack := string(debug.Stack())
				msg := fmt.Sprintf("panic: %v", rec)

				log.Printf("[PANIC][Fiber] %s %s TraceID=%v\n%s\n msg=%s ", method, path, traceID, stack, msg)

				_ = c.Status(500).SendString("Internal Server Error")
			}
		}()
		return c.Next()
	}
}
