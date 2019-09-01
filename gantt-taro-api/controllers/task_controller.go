package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/models"
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
	api.GET("/tasks", t.getAllTasks)
	api.GET("/tasks/:id", t.getTask)
	api.POST("/tasks/:id", t.updateTask)
}

func (t *TaskController) getTask(c *gin.Context) {
	id := c.GetInt("id")

	task, err := t.taskService.FindById(id)
	if err != nil {
		handleTaskError(c, err)
		return
	}

	c.JSON(200, task)
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

func (t *TaskController) updateTask(c *gin.Context) {
	var task models.Task
	err := c.Bind(&task)
	if err != nil {
		handleTaskError(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleTaskError(c, err)
		return
	}

	task.Id = id

	err = t.taskService.Update(&task)
	if err != nil {
		handleTaskError(c, err)
		return
	}

	c.Status(200)
}

func handleTaskError(c *gin.Context, err error) {
	if err == repositories.TaskNotFoundError {
		c.AbortWithStatus(404)
	} else {
		c.AbortWithStatus(500)
	}
}
