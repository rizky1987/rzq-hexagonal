package reposirory_factory

import (
	repositoryPostgresImpl "rzq-hexagonal/domain/repository/postgres"
	dbInfra "rzq-hexagonal/infrastructure/database"
)

func GeneratePostgresRepository() *RepositoryFactory {
	dbSession := dbInfra.NewPostgresDB()

	return &RepositoryFactory{
		UserRepository: repositoryPostgresImpl.NewUserRepository(dbSession),
	}
}
