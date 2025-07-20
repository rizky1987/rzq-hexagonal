package reposirory_factory

import (
	"fmt"
	"os"
	repositoryInterface "rzq-hexagonal/port/outbound"
)

type RepositoryFactory struct {
	UserRepository repositoryInterface.UserRepository
}

func NewRepositoryFactory() *RepositoryFactory {
	currentDatabaseEnv := os.Getenv("CONFIG_CURRENT_DATABASE")
	if currentDatabaseEnv == "" {
		panic(fmt.Sprintf("Please input CONFIG_CURRENT_DATABASE"))
	}

	if currentDatabaseEnv == "postgres" {
		return GeneratePostgresRepository()
	}

	panic(fmt.Sprintf("Please input CONFIG_CURRENT_DATABASE"))
}
