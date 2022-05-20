package comment

import "github.com/labstack/echo/v4"

type CommentContro interface {
	Insert() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}
