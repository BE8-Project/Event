package event

import "github.com/labstack/echo/v4"

type EventController interface {
	Insert() echo.HandlerFunc
}