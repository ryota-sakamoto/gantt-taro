package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/ryota-sakamoto/gantt-taro/models"
)

type TaskRepository interface {
	Get(int) (*models.Task, error)
	Save(*models.Task) error
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return taskRepositoryImpl{
		db: db,
	}
}

type taskRepositoryImpl struct {
	db *sqlx.DB
}

func (t taskRepositoryImpl) Get(id int) (*models.Task, error) {
	return &models.Task{}, nil
}

func (taskRepositoryImpl) Save(task *models.Task) error {
	return nil
}
