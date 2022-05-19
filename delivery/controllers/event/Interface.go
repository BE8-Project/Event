package event

import "github.com/labstack/echo/v4"

type EventController interface {
	Insert() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Get() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByUser() echo.HandlerFunc
}