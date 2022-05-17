package user

import (
	"event/delivery/helpers/request"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/entity"
	repoUser "event/repository/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type userController struct {
	Connect  repoUser.UserModel
	Validate *validator.Validate
}

func NewUserController(conn repoUser.UserModel) *userController {
	return &userController{
		Connect:  conn,
		Validate: validator.New(),
	}
}

func (u *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.InsertUser

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tipe field ada yang salah"))
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		user := entity.User{
			Name:     request.Name,
			Username: request.Username,
			Email:    request.Email,
			HP:       request.HP,
			Password: request.Password,
			Role:     2,
		}

		result, err := u.Connect.Insert(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success register User!", result))
	}
}
func (u *userController) Admin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.InsertUser

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tipe field ada yang salah"))
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		user := entity.User{
			Name:     request.Name,
			Username: request.Username,
			Email:    request.Email,
			HP:       request.HP,
			Password: request.Password,
			Role:     1,
		}

		result, err := u.Connect.Insert(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success register User!", result))
	}
}

func (u *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.Login

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tipe field ada yang salah"))
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		var input []string
		input = append(input, request.Email, request.Username, request.HP)

		login, err := u.Connect.Login(input, request.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.StatusUnauthorized(err))
		}

		if login.Token == "" {
			login.Token, _ = middlewares.CreateToken(login.ID)
		}

		return c.JSON(http.StatusOK, response.StatusOK("success login!", login))
	}
}

func (u *userController) GetUser() echo.HandlerFunc {
	return func (c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(c))
	
		result := u.Connect.GetOne(user_id)
	
		return c.JSON(http.StatusOK, response.StatusOK("success get User!", result))
	}
}

func (u *userController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(c))

		var request request.UpdateUser
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest("tipe field ada yang salah"))
		}

		user := entity.User{
			Name:     request.Name,
			Email:    request.Email,
			HP:       request.HP,
			Password: request.Password,
		}

		result, err := u.Connect.Update(&user, user_id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success update User!", result))
	}
}

func (u *userController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(c))

		result := u.Connect.Delete(user_id)

		return c.JSON(http.StatusOK, response.StatusOK("success delete User!", result))
	}
}
