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

func PostUser(context *gin.Context) {

	var input models.CreateUser

	error := context.ShouldBindJSON(&input)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	user := models.User{
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}

	connect.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
