package comment

import (
	"event/delivery/helpers/request"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/entity"
	repository "event/repository/comment"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type eventController struct {
	Connect  repository.CommentRepo
	Validate *validator.Validate
}

func NewCommentController(repo repository.CommentRepo, valid *validator.Validate) *eventController {
	return &eventController{
		Connect:  repo,
		Validate: valid,
	}
}

func (cc *eventController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(c))
		var request request.CommentInsert

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tipe field ada yang salah"))
		}

		if err := cc.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		req := entity.Comment{
			UserID:  user_id,
			Field:   request.Field,
			EventID: request.EventID,
		}
		
		strg, err := cc.Connect.Insert(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("succes comment", strg))
	}
}

func (cc *eventController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(c))
		id, _ := strconv.Atoi(c.Param("id"))

		strg, err := cc.Connect.Delete(uint(id), user_id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("delete comment", strg))
	}
}
