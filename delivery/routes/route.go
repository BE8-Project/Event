package routes

import (
	"event/delivery/controllers/category"
	"event/delivery/controllers/event"
	"event/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UsersPath(e *echo.Echo, connUser user.UserController, connCateg category.CategoruController) {
	e.POST("/register", connUser.Register())
	e.POST("/admin", connUser.Admin())
	e.POST("/login", connUser.Login())
	e.GET("/category", connCateg.GetAll())

	customer := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	customer.GET("/:username", connUser.GetUser)
	customer.PUT("/:username", connUser.Update())
	customer.DELETE("/:username", connUser.Delete())

	event := e.Group("/categories", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	event.POST("", connCateg.Insert())
	event.DELETE("/:id", connCateg.Delete())
}

func EventPath(e *echo.Echo, connect event.EventController) {
	event := e.Group("/events", middleware.JWT([]byte("$p4ssw0rd")))
	event.POST("", connect.Insert())
}
