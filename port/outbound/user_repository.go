package repositoryPostgresImpl

import "rzq-hexagonal/domain/entity"

type UserRepository interface {
	Register(user *entity.User) (*entity.User, error)
}
