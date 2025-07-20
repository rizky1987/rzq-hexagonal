package factory

import (
	"rzq-hexagonal/config"
	repositoryFactory "rzq-hexagonal/infrastructure/factory/repository"
	"rzq-hexagonal/port/inbound"
	"rzq-hexagonal/usecase"
)

type ServicesFactory struct {
	UserUseCase inbound.UserInbound
}

func NewServiceFactory() *ServicesFactory {
	//this function will read all confogiration from all environment base on
	//APP_CURRENT_ENV that we decided in application.env
	config.ConfigReader()

	repositories := repositoryFactory.NewRepositoryFactory()

	return &ServicesFactory{
		UserUseCase: usecase.NewUserUseCase(repositories),
	}
}
