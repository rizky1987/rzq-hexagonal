package router

import (
	handler "rzq-hexagonal/adapter/http/handler/fiber"
	middlewareFiber "rzq-hexagonal/cmd/http/middleware/fiber"

	"github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
	app *fiber.App
}

func NewFiberRouter() *FiberRouter {
	return &FiberRouter{app: fiber.New()}
}

func (r *FiberRouter) RegisterMiddleware() {
	r.app.Use(middlewareFiber.FiberRecover())
}

func (r *FiberRouter) RegisterRoutes(fiberHandler *handler.FiberHandler) {
	r.app.Post("/register", fiberHandler.Register)
}

func (r *FiberRouter) Start(port string) error {
	return r.app.Listen(":" + port)
}
