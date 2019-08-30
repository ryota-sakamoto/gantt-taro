package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"

	"github.com/ryota-sakamoto/gantt-taro/controllers"
	"github.com/ryota-sakamoto/gantt-taro/db"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
	"github.com/ryota-sakamoto/gantt-taro/services"
	"github.com/ryota-sakamoto/gantt-taro/utils"
)

type Server struct {
}

func (*Server) Run() error {
	r := gin.Default()
	m := melody.New()

	config := db.NewDBConfig()
	db, err := db.NewDB(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	projectRepository := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(projectRepository)
	projectController := controllers.NewProjectController(projectService)

	authenticationController := controllers.NewAuthenticationController(userService)

	r.Use(errorMiddleware)
	api := r.Group("/api/v1")
	{
		if appEnv != "development_skip_auth" {
			api.Use(validaetMiddleware)
		}

		authenticationController.AuthenticationAPI(api)
		userController.UserAPI(api)
		taskController.TaskAPI(api)
		projectController.ProjectAPI(api)
	}

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
	})

	if err := r.Run(); err != nil {
		return err
	} else {
		return nil
	}
}

func errorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) != 0 {
		c.JSON(404, gin.H{
			"message": c.Errors.JSON(),
		})
		c.Abort()
	}
}

func validaetMiddleware(c *gin.Context) {
	t := c.GetHeader("Authorization")
	claims, err := utils.ValidateToken(t)
	if err == nil {
		c.Set("user", claims)
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}
