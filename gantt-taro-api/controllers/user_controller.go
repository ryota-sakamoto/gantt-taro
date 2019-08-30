package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(p services.UserService) *UserController {
	return &UserController{
		userService: p,
	}
}

func (u *UserController) UserAPI(api *gin.RouterGroup) {
	api.GET("/users/:id", u.findByID)
}

func (u *UserController) findByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}

	user, err := u.userService.FindByID(id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, user)
}
