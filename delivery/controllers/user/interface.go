package user

import "github.com/labstack/echo/v4"

type UserController interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
	Admin() echo.HandlerFunc
}
