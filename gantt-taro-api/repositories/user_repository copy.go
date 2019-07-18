package repositories

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryota-sakamoto/gantt-taro/models"
)

type UserRepository interface {
	FindByID(int) (*models.User, error)
	Register(*models.User) error
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepositoryImpl{
		db: db,
	}
}

type userRepositoryImpl struct {
	db *sqlx.DB
}

func (u userRepositoryImpl) FindByID(id int) (*models.User, error) {
	users := []models.User{}

	err := u.db.Select(&users, "SELECT id, name FROM users where id = ?", id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("USER_NOT_FOUND")
	}

	return &users[0], nil
}

func (u userRepositoryImpl) Register(user *models.User) error {
	return nil
}
