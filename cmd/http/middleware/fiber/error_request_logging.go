package fiber_middleware

// import (
// 	"log"
// 	middlewareHelper "rzq-hexagonal/cmd/http/middleware/helper"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// func ErrorRequestLogging() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		start := time.Now()

// 		err := c.Next()
// 		status := c.Response().StatusCode()

// 		if status != fiber.StatusOK {
// 			maskedRespBody := middlewareHelper.MaskSensitiveData(c.Response().Body())
// 			headers := map[string]string{}
// 			c.Request().Header.VisitAll(func(k, v []byte) {
// 				headers[string(k)] = string(v)
// 			})

// 			maskedRequestBody := middlewareHelper.MaskSensitiveData(c.Body())
// 			traceID := c.Locals("TraceID")
// 			callerInfo := middlewareHelper.GetCallerInfo()
// 			logMsg := "[Fiber] %s %s [%d] (%v) TraceID=%v\nHeaders=%v\nBody=%s\nCaller: %s\nResponseBody %s\n"
// 			log.Printf(logMsg,
// 				c.Method(),
// 				c.Path(),
// 				status,
// 				time.Since(start),
// 				traceID,
// 				headers,
// 				maskedRequestBody,
// 				callerInfo,
// 				maskedRespBody,
// 			)
// 		}

// 		return err
// 	}
// }
