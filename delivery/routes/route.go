package routes

import (
	"event/delivery/controllers/category"
	"event/delivery/controllers/comment"
	"event/delivery/controllers/event"
	"event/delivery/controllers/participant"
	"event/delivery/controllers/user"
	"event/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func UsersPath(e *echo.Echo, connUser user.UserController) {
	e.POST("/register", connUser.Register())
	e.POST("/login", connUser.Login())

	customer := e.Group("/users", middlewares.Secret())
	customer.GET("/profile", connUser.GetUser())
	customer.PUT("/profile", connUser.Update())
	customer.DELETE("/profile", connUser.Delete())
}

func CategoryPath(e *echo.Echo, category category.CategoryController) {
	e.GET("/categories", category.GetAll())

	categories := e.Group("/admin/categories", middlewares.Secret())
	categories.POST("", category.Insert())
	categories.DELETE("/:id", category.Delete())
}

func EventPath(e *echo.Echo, connect event.EventController) {
	e.GET("/events", connect.GetAll())
	e.GET("/events/:id", connect.Get())

	event := e.Group("/admin/events", middlewares.Secret())
	event.POST("", connect.Insert())
	event.GET("", connect.GetByUser())
	event.PUT("/:id", connect.Update())
	event.DELETE("/:id", connect.Delete())
}
func CommentPath(e *echo.Echo, connect comment.CommentContro) {
	Comment := e.Group("/events/comments", middlewares.Secret())
	Comment.POST("", connect.Insert())
	Comment.DELETE("/:id", connect.Delete())
}

func ParticipantPath(e *echo.Echo, connect participant.ParticipantController) {
	participant := e.Group("/events/orders", middlewares.Secret())
	participant.POST("", connect.Insert())
	participant.GET("/:order_id/status", connect.GetStatus())
	participant.GET("/:order_id/cancel", connect.Cancel())
}
