package main

import (
	"github.com/gin-gonic/gin"
	"golang_jwt_auth/initializers"
)

func init() {
	initializers.ConnectToDb()
}

func main() {
	StartGinServer()
}

func StartGinServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
