package response

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	HP        string    `json:"hp"`
	CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Token 	 string `json:"token"`
}

type UpdateUser struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUser struct {
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}