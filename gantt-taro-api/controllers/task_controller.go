package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/services"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(t services.TaskService) *TaskController {
	return &TaskController{
		taskService: t,
	}
}

func (t *TaskController) TaskAPI(api *gin.RouterGroup) {
	api.GET("/tasks/:id", t.getTask)
	api.GET("/tasks", t.getAllTasks)
}

func (t *TaskController) getTask(c *gin.Context) {
	id := c.GetInt("id")
	if task, err := t.taskService.Get(id); err != nil {
		panic(err)
	} else {
		c.JSON(200, task)
	}
}

func (t *TaskController) getAllTasks(c *gin.Context) {
}
