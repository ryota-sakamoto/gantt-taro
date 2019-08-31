package services

import (
	"github.com/ryota-sakamoto/gantt-taro/models"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type TaskService interface {
	FindById(int) (*models.Task, error)
	GetAllTasks() ([]models.Task, error)
}

type taskServiceImpl struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(r repositories.TaskRepository) TaskService {
	return taskServiceImpl{
		taskRepository: r,
	}
}

func (t taskServiceImpl) FindById(id int) (*models.Task, error) {
	return t.taskRepository.FindById(id)
}

func (t taskServiceImpl) GetAllTasks() ([]models.Task, error) {
	return t.taskRepository.GetAllTasks()
}
