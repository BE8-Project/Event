package category

import (
	"errors"
	"event/delivery/helpers/request"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/entity"
	repository "event/repository/category"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type categController struct {
	Connect  repository.CategoryModel
	Validate *validator.Validate
}

func NewEventController(repo repository.CategoryModel) *categController {
	return &categController{
		Connect:  repo,
		Validate: validator.New(),
	}
}
func (cc *categController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_user := uint(middlewares.ExtractTokenUserId(c))

		var request request.InsertCateg
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest("data yg anda masukan salah"))
		}

		if err := cc.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		strng, err := cc.Connect.Insert(entity.Category{
			Name: request.Name,
		}, id_user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}
		return c.JSON(http.StatusCreated, response.StatusCreated(strng, request.Name))
	}
}
func (cc *categController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		strng, err := cc.Connect.Get()
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(errors.New("data tidak ditemukan")))
		}
		return c.JSON(http.StatusOK, response.StatusOK("success get Category!", strng))
	}
}
func (cc *categController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_user := uint(middlewares.ExtractTokenUserId(c))
		id_Categ, _ := strconv.Atoi(c.Param("id"))

		strng, err := cc.Connect.Delete(id_user, uint(id_Categ))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(errors.New("data tidak ditemukan")))
		}
		return c.JSON(http.StatusOK, response.StatusOK("success Delete", strng))
	}
}
