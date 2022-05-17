package routes

import (
	"event/delivery/controllers/event"
	"event/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, connUser user.UserController) {
	e.POST("/register", connUser.Register())
	e.POST("/login", connUser.Login())

	customer := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	customer.GET("/:username", connUser.GetUser)
	customer.PUT("/:username", connUser.Update())
	customer.DELETE("/:username", connUser.Delete())
}

func EventPath(e *echo.Echo, connect event.EventController) {
	event := e.Group("/events", middleware.JWT([]byte("$p4ssw0rd")))
	event.POST("", connect.Insert())
}