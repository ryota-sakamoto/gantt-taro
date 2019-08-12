package controllers

import (
	"log"
	"strconv"

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

func (u *UserController) UserAPI(api *gin.RouterGroup) {
	api.GET("/user/:id", u.findByID)
	api.POST("/user", u.register)
}

func (u *UserController) findByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}

	user, err := u.userRepository.FindByID(id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, user)
}

func (u *UserController) register(c *gin.Context) {

}
