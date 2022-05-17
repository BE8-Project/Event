package routes

import (
	"event/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, connUser user.UserController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.POST("/register", connUser.Register())
	e.POST("/login", connUser.Login())

	customer := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	customer.GET("/:username", connUser.GetUser)
	customer.PUT("/:username", connUser.Update())
	customer.DELETE("/:username", connUser.Delete())
}
