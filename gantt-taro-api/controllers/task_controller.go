package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
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
	if task, err := t.taskService.FindById(id); err != nil {
		if err == repositories.TaskNotFoundError {
			c.AbortWithStatus(404)
		} else {
			log.Println(err)
			c.AbortWithStatus(500)
		}
	} else {
		c.JSON(200, task)
	}
}

func (t *TaskController) getAllTasks(c *gin.Context) {
	tasks, err := t.taskService.GetAllTasks()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, tasks)
}
