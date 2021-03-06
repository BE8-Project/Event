package repository

import (
	"errors"
	"event/delivery/helpers/response"
	"event/entity"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *userModel {
	return &userModel{
		DB: db,
	}
}

func (u *userModel) Insert(user *entity.User) (response.User, error) {
	user.Name = strings.Title(strings.ToLower(user.Name))
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)

	record := u.DB.Create(user)

	if record.RowsAffected == 0 {
		return response.User{}, record.Error
	} else {
		return response.User{
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			HP:        user.HP,
			CreatedAt: user.CreatedAt,
		}, nil
	}
}

func (u *userModel) Login(custom []string, password string) (response.Login, error) {
	var user entity.User

	record := u.DB.Where("email = ?", custom[0]).Or("username = ?", custom[1]).Or("hp = ?", custom[2]).Find(&user)
	hash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if record.RowsAffected == 0 {
		return response.Login{}, errors.New("user or password is wrong")
	} else {
		if hash == nil {
			return response.Login{
				ID:       user.ID,
				Name:     user.Name,
			}, nil
		} else {
			return response.Login{}, errors.New("user or password is wrong")
		}
	}
}

func (u *userModel) GetOne(user_id uint) response.User {

	var user response.User
	if err := u.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		log.Warn(err)
		return response.User{}
	}
	return user
}

func (u *userModel) Delete(user_id uint) response.DeleteUser {
	var user entity.User
	if err := u.DB.Where("id = ?", user_id).Delete(&user).Error; err != nil {
		log.Warn(err)
		return response.DeleteUser{}
	}
	log.Info()
	return response.DeleteUser{
		Name:      user.Name,
		DeletedAt: user.DeletedAt,
	}
}

func (u *userModel) Update(newUser *entity.User, user_id uint) (response.UpdateUser, error) {
	var user entity.User
	if err := u.DB.Model(&user).Where("id = ?", user_id).Updates(newUser).Error; err != nil {
		log.Warn(err)
		return response.UpdateUser{}, err
	}
	log.Info()
	return response.UpdateUser{
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
