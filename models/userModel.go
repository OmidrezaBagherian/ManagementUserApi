package models

type User struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}

type CreateUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   uint8  `json:"age" binding:"required"`
}

type UpdateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}
