package handler

import (
	"rzq-hexagonal/infrastructure/factory"
)

type EchoHandler struct {
	ServicesFactory *factory.ServicesFactory
}
