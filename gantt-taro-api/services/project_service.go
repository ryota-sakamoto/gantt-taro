package services

import (
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type ProjectService interface {
}

type projectServiceImpl struct {
	projectRepository repositories.ProjectRepository
}

func NewProjectService(r repositories.ProjectRepository) ProjectService {
	return projectServiceImpl{
		projectRepository: r,
	}
}
