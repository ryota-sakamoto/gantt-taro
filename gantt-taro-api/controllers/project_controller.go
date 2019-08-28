package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/services"
)

type ProjectController struct {
	projectService services.ProjectService
}

func NewProjectController(p services.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: p,
	}
}

func (p ProjectController) ProjectAPI(api *gin.RouterGroup) {

}
