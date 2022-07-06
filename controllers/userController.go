package controllers

import (
	"UserApi/connect"
	"UserApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(context *gin.Context) {
	var users []models.User
	connect.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}
