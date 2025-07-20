package repository

import (
	"rzq-hexagonal/domain/entity"
	outboundInterface "rzq-hexagonal/port/outbound"

	"gorm.io/gorm"
)

type UserRepositoryPostgreImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) outboundInterface.UserRepository {
	return &UserRepositoryPostgreImpl{db: db}
}

func (r *UserRepositoryPostgreImpl) Register(user *entity.User) (*entity.User, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
