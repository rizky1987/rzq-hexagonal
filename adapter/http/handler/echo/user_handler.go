package handler

import (
	"net/http"
	"rzq-hexagonal/adapter/http/request"
	"rzq-hexagonal/adapter/http/response"
	"rzq-hexagonal/infrastructure/factory"
    
	"github.com/labstack/echo/v4"
	"fmt"
)

type EchoHandler struct {
	ServicesFactory *factory.ServicesFactory
}

func (h *EchoHandler) Register(c echo.Context) error {

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	var req request.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.ServicesFactory.UserUseCase.Register(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	result := response.ConvertUserFromEntity(user)

	return c.JSON(http.StatusOK, result)
}
