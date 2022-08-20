package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type EchoRequest struct {
	Msg string `json:"message"`
}

type EchoRespond struct {
	Msg string `json:"message"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/echo", func(c *gin.Context) {
		request := EchoRequest{}
		if err := c.BindJSON(&request); err != nil {
			fmt.Printf("Error: %v", err)
			c.Abort()
			return
		}

		time.Sleep(time.Duration((300 + rand.Intn(200))) * time.Millisecond)
		c.JSON(200, &EchoRespond{
			Msg: "respond:" + request.Msg,
		})
	})
	r.Run("0.0.0.0:8888")
}
