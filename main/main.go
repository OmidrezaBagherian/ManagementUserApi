package main

import (
	"UserApi/connect"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	connect.ConnectDatabase()

	route.Run()
}
