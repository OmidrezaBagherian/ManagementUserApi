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

	route.Run()
}
