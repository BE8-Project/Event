package category

import "github.com/labstack/echo/v4"

type CategoryController interface {
	Insert() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
