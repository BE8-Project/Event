package participant

import (
	"encoding/json"
	"errors"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/entity"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	token string
)

func TestCreateToken(t *testing.T) {
	token, _ = middlewares.CreateToken(1)
}

func TestRegister(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"event_id": 1,
			"payment_type" : "gopay",
			"total" : 10000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewParticipantController(&mockOrder{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success create Order!", resp.Message)
		assert.Equal(t,  map[string]interface {}{"created_at":"0001-01-01T00:00:00Z", "order_id":"", "payment_simulator":interface {}(nil), "payment_type":"gopay", "payment_url":"https://api.sandbox.midtrans.com/v2/gopay//qr-code", "status":"pending", "total":float64(0)}, resp.Data)
	})

	t.Run("Status BadRequest Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"event_id": "1",
			"payment_type" : "gopay",
			"total" : 10000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewParticipantController(&mockOrder{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "field=event_id, expected=string", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status BadRequest Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"event_id": 1,
			"payment_type" : "",
			"total" : 10000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewParticipantController(&mockOrder{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message []string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, []string([]string{"field PaymentType : required"}), resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status BadRequest Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"event_id": 1,
			"payment_type" : "gopay",
			"total" : 10000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewParticipantController(&mockError{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "error", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

type mockOrder struct {}

func (m *mockOrder) Insert(participant *entity.Participant) (response.InsertParticipant, error) {
	return response.InsertParticipant{}, nil
}

func (m *mockOrder) Update(order_id string, user_id uint, participant *entity.Participant) (response.UpdateParticipat, error) {
	return response.UpdateParticipat{}, nil
}

func (m *mockOrder) GetByUser(user_id uint) []response.GetParticipant {
	return []response.GetParticipant{}
}

type mockError struct {}

func (m *mockError) Insert(participant *entity.Participant) (response.InsertParticipant, error) {
	return response.InsertParticipant{}, errors.New("error")
}

func (m *mockError) Update(order_id string, user_id uint, participant *entity.Participant) (response.UpdateParticipat, error) {
	return response.UpdateParticipat{}, nil
}

func (m *mockError) GetByUser(user_id uint) []response.GetParticipant {
	return []response.GetParticipant{}
}