package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Style struct {
	Width  int64
	Height int64
}

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var style Style
		if err := json.Unmarshal(msg, &style); err == nil {
			m.Broadcast(msg)
		}
	})

	r.Run()
}
