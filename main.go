package main

import (
	"event/config"
	"event/delivery/controllers/event"
	"event/delivery/controllers/user"
	"event/delivery/middlewares"
	"event/delivery/routes"
	eventModel "event/repository/event"
	userModel "event/repository/user"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.AutoMigrate(db)

	e := echo.New()

	userModel := userModel.NewUserModel(db)
	userController := user.NewUserController(userModel)

	middlewares.General(e)
	routes.Route(e, userController)
	routes.EventPath(e, event.NewEventController(eventModel.NewEventModel(db), validator.New()))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
