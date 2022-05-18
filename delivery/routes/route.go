package routes

import (
	"event/delivery/controllers/category"
	"event/delivery/controllers/event"
	"event/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UsersPath(e *echo.Echo, connUser user.UserController) {
	e.POST("/register", connUser.Register())
	e.POST("/login", connUser.Login())

	customer := e.Group("/profile", middleware.JWT([]byte("$p4ssw0rd")))
	customer.GET("", connUser.GetUser())
	customer.PUT("", connUser.Update())
	customer.DELETE("", connUser.Delete())
}

func CategoryPath(e *echo.Echo, category category.CategoryController) {
	e.GET("/categories", category.GetAll())

	categories := e.Group("/admin/categories", middleware.JWT([]byte("$p4ssw0rd")))
	categories.POST("", category.Insert())
	categories.DELETE("/:id", category.Delete())
}

func EventPath(e *echo.Echo, connect event.EventController) {
	e.GET("/events", connect.GetAll())

	event := e.Group("/admin/events", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	event.POST("", connect.Insert())
}
