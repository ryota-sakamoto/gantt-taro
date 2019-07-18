package repositories

import "github.com/jmoiron/sqlx"

type UserRepository interface {
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepositoryImpl{
		db: db,
	}
}

type userRepositoryImpl struct {
	db *sqlx.DB
}
