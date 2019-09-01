package repositories

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryota-sakamoto/gantt-taro/models"
)

var (
	TaskNotFoundError = errors.New("task is not found")
)

type TaskRepository interface {
	FindById(int) (*models.Task, error)
	GetAllTasks() ([]models.Task, error)
	Save(*models.Task) error
	Update(*models.Task) error
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return taskRepositoryImpl{
		db: db,
	}
}

type taskRepositoryImpl struct {
	db *sqlx.DB
}

func (t taskRepositoryImpl) FindById(id int) (*models.Task, error) {
	tasks := []models.Task{}
	err := t.db.Select(&tasks, "SELECT id, name, started_at, ended_at FROM tasks WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, TaskNotFoundError
	}

	return &tasks[0], nil
}

func (t taskRepositoryImpl) GetAllTasks() ([]models.Task, error) {
	tasks := []models.Task{}
	err := t.db.Select(&tasks, "SELECT id, name, started_at, ended_at FROM tasks")
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (taskRepositoryImpl) Save(task *models.Task) error {
	return nil
}

func (t taskRepositoryImpl) Update(task *models.Task) error {
	_, err := t.db.Exec("UPDATE tasks SET name = ?, started_at = ?, ended_at = ? WHERE id = ?",
		task.Name, task.StartedAt.ToTime(), task.EndedAt.ToTime(), task.Id)

	return err
}
