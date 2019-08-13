package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"

	"github.com/ryota-sakamoto/gantt-taro/controllers"
	"github.com/ryota-sakamoto/gantt-taro/db"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
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

	user_repository := repositories.NewUserRepository(db)
	user_controller := controllers.NewUserController(user_repository)

	task_repository := repositories.NewTaskRepository(db)
	task_controller := controllers.NewTaskController(task_repository)

	project_repository := repositories.NewProjectRepository(db)
	project_controller := controllers.NewProjectController(project_repository)

	r.Use(errorMiddleware)
	api := r.Group("/api")
	{
		api.Use(validaetMiddleware)

		user_controller.UserAPI(api)
		task_controller.TaskAPI(api)
		project_controller.ProjectAPI(api)
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
	if t != "" && utils.ValidateToken(t) {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}
