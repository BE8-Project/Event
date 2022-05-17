package event

import (
	"event/delivery/helpers/request"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/entity"
	repository "event/repository/event"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type eventController struct {
	Connect repository.EventModel
	Validate *validator.Validate
}

func NewEventController(repo repository.EventModel, valid *validator.Validate) *eventController {
	return &eventController{
		Connect: repo,
		Validate : valid,
	}
}

func (c *eventController) Insert() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		username := middlewares.ExtractTokenUsername(ctx)
		var request request.InsertEvent

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := c.Validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		event := entity.Event{
			Name: request.Name,
			HostedBy: request.HostedBy,
			Date: request.Date,
			Location: request.Location,
			Details: request.Details,
			Ticket: request.Ticket,
			CategoryID: request.CategoryID,
			Username: username,
		}

		result := c.Connect.Insert(&event)

		return ctx.JSON(http.StatusCreated, response.StatusCreated("Created", result))
	}
}