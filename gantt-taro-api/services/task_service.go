package services

import (
	"github.com/ryota-sakamoto/gantt-taro/models"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type TaskService interface {
	Get(int) (*models.Task, error)
}

type taskServiceImpl struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(r repositories.TaskRepository) TaskService {
	return taskServiceImpl{
		taskRepository: r,
	}
}

func (t taskServiceImpl) Get(id int) (*models.Task, error) {
	return t.taskRepository.Get(id)
}
