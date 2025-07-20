package handler

import (
	"rzq-hexagonal/adapter/http/request"
	"rzq-hexagonal/adapter/http/response"
	"rzq-hexagonal/infrastructure/factory"

	"github.com/gofiber/fiber/v2"
)

type FiberHandler struct {
	ServicesFactory *factory.ServicesFactory
}

func (h *FiberHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	user, err := h.ServicesFactory.UserUseCase.Register(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	result := response.ConvertUserFromEntity(user)
	return c.JSON(result)
}
