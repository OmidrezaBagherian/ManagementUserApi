package main

import (
	"UserApi/connect"
	"UserApi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	connect.ConnectDatabase()

	route.GET("/users", controllers.GetUser)
	route.POST("/users", controllers.PostUser)
	route.GET("/users/:id", controllers.FindUser)

	route.Run()
}
