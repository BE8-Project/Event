package main

import (
	"event/config"
	"event/delivery/controllers/user"
	"event/delivery/routes"
	userModel "event/repository/user"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.AutoMigrate(db)

	e := echo.New()

	userModel := userModel.NewUserModel(db)
	userController := user.NewUserController(userModel)

	routes.Route(e, userController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
