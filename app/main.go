package main

import (
	"github.com/gin-gonic/gin"
	"golang_jwt_auth/app/controllers"
	"golang_jwt_auth/app/initializers"
)

func init() {
	initializers.ConnectToDb()
	initializers.LoadEnv()
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
