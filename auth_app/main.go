package main

import (
	"auth_app/controllers"
	"auth_app/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	StartGinServer()
}

func StartGinServer() {
	r := gin.Default()
	r.POST("/users/access", controllers.GenerateTokens)
	r.POST("/users/refresh", controllers.RefreshTokens)

	err := r.Run()

	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
