package main

import (
	"event/config"
	"event/delivery/controllers/category"
	"event/delivery/controllers/comment"
	"event/delivery/controllers/event"
	"event/delivery/controllers/participant"
	"event/delivery/controllers/user"
	"event/delivery/middlewares"
	"event/delivery/routes"
	categModel "event/repository/category"
	eventModel "event/repository/event"
	participantModel "event/repository/participant"
	userModel "event/repository/user"
	commentModel "event/repository/comment"
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

	categModel := categModel.NewCategotyModel(db)
	categController := category.NewEventController(categModel)

	middlewares.General(e)
	routes.UsersPath(e, userController)
	routes.CategoryPath(e, categController)
	routes.EventPath(e, event.NewEventController(eventModel.NewEventModel(db), validator.New()))
	routes.ParticipantPath(e, participant.NewParticipantController(participantModel.NewParticipantModel(db), validator.New()))
	routes.CommentPath(e,comment.NewCommentController(commentModel.NewCommenttModel(db),validator.New()))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
