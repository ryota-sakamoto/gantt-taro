package services

import (
	"github.com/ryota-sakamoto/gantt-taro/models"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type UserService interface {
	Register(*models.User) error
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

func (u userServiceImpl) Register(model *models.User) error {
	user, err := u.userRepository.FindByUniqiueId(model.UniqueId)
	if err != nil && err != repositories.UserNotFoundError {
		return err
	}

	if user != nil {
		return nil
	}

	return u.userRepository.Register(model)
}
