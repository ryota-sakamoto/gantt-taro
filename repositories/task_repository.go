package repositories

import (
	"github.com/ryota-sakamoto/gantt-taro/models"
)

type TaskRepository interface {
	Get(int) (*models.Task, error)
	Save(*models.Task) error
}

type TaskRepositoryImpl struct {
}

func (TaskRepositoryImpl) Get(id int) (*models.Task, error) {
	return &models.Task{}, nil
}

func (TaskRepositoryImpl) Save(task *models.Task) error {
	return nil
}
