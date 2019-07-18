package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type UserController struct {
	userRepository repositories.UserRepository
}

func NewUserController(p repositories.UserRepository) *UserController {
	return &UserController{
		userRepository: p,
	}
}

func (p UserController) UserAPI(api *gin.RouterGroup) {

}
