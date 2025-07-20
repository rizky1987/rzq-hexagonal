package reposirory_factory

import (
	"fmt"
	"os"
	repositoryPostgresInterface "rzq-hexagonal/port/outbound/repository/postgres"
)

type RepositoryFactory struct {
	UserRepository repositoryPostgresInterface.UserRepository
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
