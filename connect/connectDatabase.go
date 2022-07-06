package connect

import (
	"UserApi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	var database, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("Fail to connect to database!!!")
	}
	database.AutoMigrate(&models.User{})

	DB = database
}
