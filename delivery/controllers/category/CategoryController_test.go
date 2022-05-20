package category

import (
	"encoding/json"
	"errors"
	"event/delivery/middlewares"
	"event/entity"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var token_admin string

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token_admin, _ = middlewares.CreateToken(1)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "music",
		})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewEventController(&mockCateg{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "create category Success", resp.Message)
		assert.Equal(t, "music", resp.Data)
	})
	t.Run("Status Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "music",
		})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewEventController(&mockErrorCategInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "gagal input", resp.Message)
	})
	t.Run("StatusInvalidRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": 55555,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewEventController(&mockErrorCategInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "data yg anda masukan salah", resp.Message)
	})
	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nomor": 55555,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewEventController(&mockErrorCategInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "", resp.Message)
	})

}
func TestGetAll(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewEventController(&mockCateg{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get Category!", resp.Message)
		assert.Equal(t, []interface{}([]interface{}{map[string]interface{}{"CreatedAt": "0001-01-01T00:00:00Z", "DeletedAt": interface{}(nil), "Events": interface{}(nil), "ID": float64(0), "Name": "music", "UpdatedAt": "0001-01-01T00:00:00Z"}}), resp.Data)
	})
	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewEventController(&mockErrorCategInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "data tidak ditemukan", resp.Message)
	})
}
func TestDelete(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewEventController(&mockCateg{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success Delete", resp.Message)
		assert.Equal(t, "success delete Category", resp.Data)
	})
	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewEventController(&mockErrorCategInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "data tidak ditemukan", resp.Message)
	})

}

type mockCateg struct{}

func (u *mockCateg) Insert(categ entity.Category, id uint) (string, error) {
	return "create category Success", nil
}

func (u *mockCateg) Get() ([]entity.Category, error) {
	return []entity.Category{
		{
			Name: "music",
		},
	}, nil
}
func (u *mockCateg) Delete(id_user, id_categ uint) (string, error) {
	return "success delete Category", nil
}

type mockErrorCategInput struct{}

func (u *mockErrorCategInput) Insert(categ entity.Category, id uint) (string, error) {
	return "", errors.New("gagal input")
}

func (u *mockErrorCategInput) Get() ([]entity.Category, error) {
	return []entity.Category{}, errors.New("data tidak ditemukan")
}
func (u *mockErrorCategInput) Delete(id_user, id_categ uint) (string, error) {
	return "", errors.New("data tidak ditemukan")
}
