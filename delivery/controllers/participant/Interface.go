package participant

import "github.com/labstack/echo/v4"

type ParticipantController interface {
	Insert() echo.HandlerFunc
	GetStatus() echo.HandlerFunc
	Cancel() echo.HandlerFunc
}