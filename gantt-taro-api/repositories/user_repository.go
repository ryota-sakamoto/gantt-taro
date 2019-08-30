package repositories

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryota-sakamoto/gantt-taro/models"
)

var (
	UserNotFoundError = errors.New("user is not found")
)

type UserRepository interface {
	FindByID(int) (*models.User, error)
	FindByUniqiueId(string) (*models.User, error)
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
	return u.findUserWithQuery("SELECT id, unique_id FROM users where id = ?", id)
}

func (u userRepositoryImpl) FindByUniqiueId(id string) (*models.User, error) {
	return u.findUserWithQuery("SELECT id, unique_id FROM users where unique_id = ?", id)
}

func (u userRepositoryImpl) Register(user *models.User) error {
	_, err := u.db.Exec("INSERT INTO users(unique_id) VALUE(?)", user.UniqueId)
	if err != nil {
		return err
	}
	return nil
}

func (u userRepositoryImpl) findUserWithQuery(query string, args ...interface{}) (*models.User, error) {
	users := []models.User{}

	err := u.db.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, UserNotFoundError
	}

	return &users[0], nil
}
