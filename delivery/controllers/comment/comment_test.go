package comment

import (
	"encoding/json"
	"errors"
	"event/delivery/middlewares"
	"event/entity"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
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
			"eventId": 1,
			"field":   "music",
		})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCommentController(&mockComment{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "succes comment", resp.Message)
		assert.Equal(t, "Comment success", resp.Data)
	})
	t.Run("Status BIND", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"field": 57688,
		})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCommentController(&mockComment{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "tipe field ada yang salah", resp.Message)
	})
	t.Run("Status validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCommentController(&mockComment{}, validator.New())

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
	t.Run("Status badrequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"eventId": 1,
			"field":   "music",
		})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCommentController(&mockCommentEror{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
	})

}
func TestDelete(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"eventId": 1,
		})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/comments/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewCommentController(&mockComment{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "delete comment", resp.Message)
		assert.Equal(t, "delete comment", resp.Data)
	})
	t.Run("Status badrequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"eventId": 1,
			"field":   "music",
		})

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events/comments/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewCommentController(&mockCommentEror{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)

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
		context.SetPath("/comments/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewCommentController(&mockComment{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "get comment", resp.Message)
		assert.Equal(t, []interface{}([]interface{}{}), resp.Data)
	})
	t.Run("Status badrequest", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		context.SetPath("/comments/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewCommentController(&mockCommentEror{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
	})
}

type mockComment struct{}

func (u *mockComment) Insert(comment *entity.Comment) (string, error) {
	return "Comment success", nil
}
func (u *mockComment) Get(id uint) ([]entity.Comment, error) {
	return []entity.Comment{}, nil
}
func (u *mockComment) Delete(id uint, idUser uint) (string, error) {
	return "delete comment", nil
}

type mockCommentEror struct{}

func (u *mockCommentEror) Insert(comment *entity.Comment) (string, error) {
	return "", errors.New("")
}
func (u *mockCommentEror) Get(id uint) ([]entity.Comment, error) {
	return []entity.Comment{}, errors.New("")
}
func (u *mockCommentEror) Delete(id uint, idUser uint) (string, error) {
	return "", errors.New("")
}
