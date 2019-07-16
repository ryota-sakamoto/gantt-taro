package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"

	"github.com/ryota-sakamoto/gantt-taro/config"
	"github.com/ryota-sakamoto/gantt-taro/controllers"
	"github.com/ryota-sakamoto/gantt-taro/repositories"
)

type Server struct {
}

func (*Server) Run() error {
	r := gin.Default()
	m := melody.New()

	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	task_repository := repositories.NewTaskRepository(db)
	task_controller := controllers.NewTaskController(task_repository)

	api := r.Group("/api")
	{
		task_controller.TaskAPI(api)
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
