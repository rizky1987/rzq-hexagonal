package router

import (
	handler "rzq-hexagonal/adapter/http/handler/echo"
	middlewareEcho "rzq-hexagonal/cmd/http/middleware/echo"

	"github.com/labstack/echo/v4"
)

type EchoRouter struct {
	app *echo.Echo
}

func NewEchoRouter() *EchoRouter {
	return &EchoRouter{app: echo.New()}
}

func (r *EchoRouter) RegisterMiddleware() {
	r.app.Use(middlewareEcho.TrackerMiddleware)
	r.app.Use(middlewareEcho.EchoRateLimiter())
	//r.app.Use(middlewareEcho.ErrorRequestLogging())
	r.app.Use(middlewareEcho.EchoRecover())
}

func (r *EchoRouter) RegisterRoutes(echoHandler *handler.EchoHandler) {
}

func (r *EchoRouter) Start(port string) error {
	return r.app.Start(":" + port)
}
