package inbound

import (
	"rzq-hexagonal/adapter/http/request"
	"rzq-hexagonal/domain/entity"
)

type UserInbound interface {
	Register(req request.RegisterRequest) (*entity.User, error)
}
