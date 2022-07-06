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

func FindUser(context *gin.Context) {
	var user models.User
	error := connect.DB.Where("id=?", context.Param("id")).First(&user).Error
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Not Found",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func UpdateUser(context *gin.Context) {
	var user models.User
	error := connect.DB.Where("id=?", context.Param("id")).First(&user).Error
	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Not Found",
		})
		return
	}

	var input models.UpdateUser
	errorInput := context.ShouldBindJSON(&input)
	if errorInput != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorInput.Error()})
		return
	}

	updateUser := models.User{
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}

	connect.DB.Model(&user).Updates(&updateUser)
	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
