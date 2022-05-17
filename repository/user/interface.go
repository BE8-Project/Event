package repository

import (
	"event/delivery/helpers/response"
	"event/entity"
)

type UserModel interface {
	Insert(user *entity.User) (response.User, error)
	Login(custom []string, password string) (response.Login, error)
	GetOne(user_id uint) response.User
	Delete(user_id uint) response.DeleteUser
	Update(newUser *entity.User, user_id uint) (response.UpdateUser, error)
}
