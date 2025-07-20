package usecase

import (
	"errors"
	"rzq-hexagonal/adapter/http/request"
	"rzq-hexagonal/domain/entity"
	repositoryFactory "rzq-hexagonal/infrastructure/factory/repository"
	"rzq-hexagonal/port/inbound"

	"github.com/go-playground/validator/v10"
)

type UserUseCase struct {
	RepositoryFactory *repositoryFactory.RepositoryFactory
	Validate          *validator.Validate
}

func NewUserUseCase(repositoryFactory *repositoryFactory.RepositoryFactory) inbound.UserInbound {
	return &UserUseCase{
		RepositoryFactory: repositoryFactory,
		Validate:          validator.New(),
	}
}

func (uc *UserUseCase) Register(req request.RegisterRequest) (*entity.User, error) {

	if err := uc.Validate.Struct(req); err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	userEntity := &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := uc.RepositoryFactory.UserRepository.Register(userEntity)
	if err != nil || user == nil {
		return nil, err
	}

	return user, nil
}
