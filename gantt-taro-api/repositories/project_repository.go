package repositories

import "github.com/jmoiron/sqlx"

type ProjectRepository interface {
}

func NewProjectRepository(db *sqlx.DB) ProjectRepository {
	return projectRepositoryImpl{
		db: db,
	}
}

type projectRepositoryImpl struct {
	db *sqlx.DB
}
