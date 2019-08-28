package services

import (
	"github.com/ryota-sakamoto/gantt-taro/models"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type UserService interface {
	FindByID(int) (*models.User, error)
}

type userServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return userServiceImpl{
		userRepository: r,
	}
}

func (u userServiceImpl) FindByID(id int) (*models.User, error) {
	return u.FindByID(id)
}
