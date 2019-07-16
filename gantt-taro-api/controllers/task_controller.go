package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type TaskController struct {
	taskRepository repositories.TaskRepository
}

func NewTaskController(t repositories.TaskRepository) *TaskController {
	return &TaskController{
		taskRepository: t,
	}
}

func (t *TaskController) TaskAPI(api *gin.RouterGroup) {
	api.GET("/task/:id", t.getTask)
	api.GET("/tasks", t.getAllTasks)
}

func (t *TaskController) getTask(c *gin.Context) {
	id := c.GetInt("id")
	if task, err := t.taskRepository.Get(id); err != nil {
		panic(err)
	} else {
		c.JSON(200, task)
	}
}

func (t *TaskController) getAllTasks(c *gin.Context) {
}
