package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type ProjectController struct {
	projectRepository repositories.ProjectRepository
}

func NewProjectController(p repositories.ProjectRepository) *ProjectController {
	return &ProjectController{
		projectRepository: p,
	}
}

func (p ProjectController) ProjectAPI(api *gin.RouterGroup) {

}
