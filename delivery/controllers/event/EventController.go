package event

import (
	"event/delivery/helpers/request"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/delivery/usecase"
	"event/entity"
	repository "event/repository/event"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type eventController struct {
	Connect  repository.EventModel
	Validate *validator.Validate
}

func NewEventController(repo repository.EventModel, valid *validator.Validate) *eventController {
	return &eventController{
		Connect:  repo,
		Validate: valid,
	}
}

func (c *eventController) Insert() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(ctx))
		var request request.InsertEvent

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tipe field ada yang salah"))
		}
		file, err := ctx.FormFile("file")
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tidak ada file yang diupload"))
		}
		img, err := usecase.Upload(file)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestUpload(err))
		}
		if err := c.Validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		start, end := request.DateStart, request.DateEnd
		layout := "2006-01-02T15:04"
		convertStart, err := time.Parse(layout, start)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("input date start salah"))
		}

		convertEnd, err := time.Parse(layout, end)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("input date end salah"))
		}

		event := entity.Event{
			Name:       request.Name,
			HostedBy:   request.HostedBy,
			DateStart:  convertStart,
			DateEnd:    convertEnd,
			Location:   request.Location,
			Image:      img,
			Details:    request.Details,
			Ticket:     request.Ticket,
			Price: 		request.Price,
			CategoryID: request.CategoryID,
			UserID:     user_id,
		}

		result := c.Connect.Insert(&event)

		return ctx.JSON(http.StatusCreated, response.StatusCreated("Berhasil membuat Event!", result))
	}
}

func (c *eventController) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		location := ctx.QueryParam("location")
		limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
		page, _ := strconv.Atoi(ctx.QueryParam("page"))

		results := c.Connect.GetAll(name, location, limit, page)

		if len(results) == 0 {
			return ctx.JSON(http.StatusNotFound, response.StatusNotFound("Data tidak ditemukan!"))
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("Berhasil mengambil semua Event!", results))
	}
}

func (c *eventController) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, _ := strconv.Atoi(ctx.Param("id"))

		result, err := c.Connect.Get(uint(id))

		if err != nil {
			return ctx.JSON(http.StatusNotFound, response.StatusNotFound("Data tidak ditemukan!"))
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("Berhasil mengambil Event!", result))
	}
}

func (c *eventController) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var event entity.Event
		var request request.UpdateEvent
		var requestStart, requestEnd time.Time

		user_id := uint(middlewares.ExtractTokenUserId(ctx))
		id, _ := strconv.Atoi(ctx.Param("id"))

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tipe field ada yang salah"))
		}

		start, end := request.DateStart, request.DateEnd
		if request.DateStart != "" && request.DateEnd != "" {
			layout := "2006-01-02T15:04"

			convertStart, err := time.Parse(layout, start)

			if err != nil {
				return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("input date start salah"))
			}

			convertEnd, err := time.Parse(layout, end)

			if err != nil {
				return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("input date end salah"))
			}

			requestStart, requestEnd = convertStart, convertEnd
		}

		event = entity.Event{
			Name:       request.Name,
			HostedBy:   request.HostedBy,
			DateStart:  requestStart,
			DateEnd:    requestEnd,
			Location:   request.Location,
			Details:    request.Details,
			Ticket:     request.Ticket,
			Price: 		request.Price,
			CategoryID: request.CategoryID,
		}

		result, err := c.Connect.Update(uint(id), user_id, &event)

		if err != nil {
			if err.Error() == "required" {
				return ctx.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tidak ada field yang dimasukkan"))
			} else {
				return ctx.JSON(http.StatusForbidden, response.StatusForbidden(err.Error()))
			}
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("Berhasil mengupdate Event!", result))
	}
}

func (c *eventController) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(ctx))
		id, _ := strconv.Atoi(ctx.Param("id"))

		result, err := c.Connect.Delete(uint(id), user_id)

		if err != nil {
			return ctx.JSON(http.StatusForbidden, response.StatusForbidden(err.Error()))
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("Berhasil menghapus Event!", result))
	}
}

func (c *eventController) GetByUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(ctx))

		results := c.Connect.GetByUser(user_id)

		if len(results) == 0 {
			return ctx.JSON(http.StatusNotFound, response.StatusNotFound("Data tidak ditemukan!"))
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("Berhasil mengambil Event!", results))
	}
}
