package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ryota-sakamoto/gantt-taro/models"
	"github.com/ryota-sakamoto/gantt-taro/services"
	"github.com/ryota-sakamoto/gantt-taro/utils"
)

type AuthenticationController struct {
	userService services.UserService
}

func NewAuthenticationController(p services.UserService) *AuthenticationController {
	return &AuthenticationController{
		userService: p,
	}
}

func (a *AuthenticationController) AuthenticationAPI(api *gin.RouterGroup) {
	api.POST("/login", a.login)
	api.POST("/logout", a.logout)
}

func (a *AuthenticationController) login(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		return
	}
	claim := u.(*utils.Claims)
	user := models.User{
		UniqueId: claim.UniqueId,
	}

	err := a.userService.Register(&user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
	}
}

func (a *AuthenticationController) logout(c *gin.Context) {

}
