package event

import (
	"encoding/json"
	"errors"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/entity"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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
			"name": "status completed",
			"hosted_by": "Murni Sekali",
			"date_start": "2022-12-17T19:47",
			"date_end": "2022-12-17T19:47",
			"location": "jakarta",
			"details": "presmian dunia digital",
			"ticket": 5,
			"category_id" : 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "Berhasil membuat Event!", resp.Message)
		assert.Equal(t,  map[string]interface {}{"created_at":"0001-01-01T00:00:00Z", "name":"status completed"}, resp.Data)
	})

	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": 1,
			"hosted_by": "Murni Sekali",
			"date_start": "2022-12-17T19:47",
			"date_end": "2022-12-17T19:47",
			"location": "jakarta",
			"details": "presmian dunia digital",
			"ticket": 5,
			"category_id" : 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "tipe field ada yang salah", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Required", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "",
			"hosted_by": "Murni Sekali",
			"date_start": "2022-12-17T19:47",
			"date_end": "2022-12-17T19:47",
			"location": "jakarta",
			"details": "presmian dunia digital",
			"ticket": 5,
			"category_id" : 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message []string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, []string{"field Name : required"}, resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Input Date Start", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "Hosting",
			"hosted_by": "Murni Sekali",
			"date_start": "2022-12-17T19:47:00",
			"date_end": "2022-12-17T19:47",
			"location": "jakarta",
			"details": "presmian dunia digital",
			"ticket": 5,
			"category_id" : 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "input date start salah", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Input Date End", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "Hosting",
			"hosted_by": "Murni Sekali",
			"date_start": "2022-12-17T19:47",
			"date_end": "2022-12-17T19:47:35",
			"location": "jakarta",
			"details": "presmian dunia digital",
			"ticket": 5,
			"category_id" : 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "input date end salah", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")
		controller := NewEventController(&mockEvent{}, validator.New())
		controller.GetAll()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "Berhasil mengambil semua Event!", resp.Message)
		assert.Equal(t, []interface {}([]interface {}{map[string]interface {}{"date_end":"0001-01-01T00:00:00Z", "date_start":"0001-01-01T00:00:00Z", "details":"presmian dunia digital", "hosted_by":"Murni Sekali", "id":float64(1), "location":"jakarta", "name":"Webinar sekali setahun", "ticket":float64(5)}, map[string]interface {}{"date_end":"0001-01-01T00:00:00Z", "date_start":"0001-01-01T00:00:00Z", "details":"presmian dunia digital", "hosted_by":"Murni Sekali", "id":float64(2), "location":"jakarta", "name":"Testing ke 3", "ticket":float64(5)}}), resp.Data)
	})

	t.Run("Status Notfound", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")
		controller := NewEventController(&mockError{}, validator.New())
		controller.GetAll()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data tidak ditemukan!", resp.Message)
		assert.Nil(t, resp.Data)	
	})
}

func TestGet(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")
		controller := NewEventController(&mockEvent{}, validator.New())
		controller.Get()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "Berhasil mengambil Event!", resp.Message)
		assert.Equal(t, map[string]interface {}(map[string]interface {}{"date_end":"0001-01-01T00:00:00Z", "date_start":"0001-01-01T00:00:00Z", "details":"presmian dunia digital", "hosted_by":"Murni Sekali", "id":float64(1), "location":"jakarta", "name":"Webinar sekali setahun", "ticket":float64(5)}), resp.Data)
	})

	t.Run("Status Notfound", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/events")
		controller := NewEventController(&mockError{}, validator.New())
		controller.Get()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		
		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Data tidak ditemukan!", resp.Message)
		assert.Nil(t, resp.Data)	
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "status completed",
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "Berhasil mengupdate Event!", resp.Message)
		assert.Equal(t,  map[string]interface {}{"name":"status completed", "updated_at":"0001-01-01T00:00:00Z"}, resp.Data)
	})

	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "tipe field ada yang salah", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Required", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockError{}, validator.New())
		middlewares.Secret()(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		fmt.Println(res.Body.String())

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "tidak ada field yang dimasukkan", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockErrorForbidden{}, validator.New())
		middlewares.Secret()(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "you are not allowed to access this resource", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Input Date Start", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "Hosting",
			"hosted_by": "Murni Sekali",
			"date_start": "2022-12-17T19:47:00",
			"date_end": "2022-12-17T19:47",
			"location": "jakarta",
			"details": "presmian dunia digital",
			"ticket": 5,
			"category_id" : 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "input date start salah", resp.Message)
		assert.Nil(t, resp.Data)
	})

	t.Run("Status Input Date End", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "Hosting",
			"hosted_by": "Murni Sekali",
			"date_start": "2022-12-17T19:47",
			"date_end": "2022-12-17T19:47:35",
			"location": "jakarta",
			"details": "presmian dunia digital",
			"ticket": 5,
			"category_id" : 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "input date end salah", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockEvent{}, validator.New())
		middlewares.Secret()(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "Berhasil menghapus Event!", resp.Message)
		assert.Equal(t, map[string]interface {}{"deleted_at":interface {}(nil), "name":"event 1"}, resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/events/1")
		controller := NewEventController(&mockErrorForbidden{}, validator.New())
		middlewares.Secret()(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "you are not allowed to access this resource", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

type mockEvent struct {}

func (m *mockEvent) Insert(task *entity.Event) response.InsertEvent {
	return response.InsertEvent{
		Name: task.Name,
		CreatedAt: task.CreatedAt,
	}
}

func (m *mockEvent) GetAll(name, location string, limit, page int) []response.GetEvent {
	return []response.GetEvent{
		{
			ID: 1,
			Name: "Webinar sekali setahun",
			HostedBy: "Murni Sekali",
			DateStart: time.Time{},
			DateEnd: time.Time{},
			Location: "jakarta",
			Details: "presmian dunia digital",
			Ticket: 5,
		},
		{
			ID: 2,
			Name: "Testing ke 3",
			HostedBy: "Murni Sekali",
			DateStart:  time.Time{},
			DateEnd:  time.Time{},
			Location: "jakarta",
			Details: "presmian dunia digital",
			Ticket: 5,
		},
	}
}

func (m *mockEvent) Get(id uint) (response.GetEvent, error) {
	return response.GetEvent{
		ID: 1,
		Name: "Webinar sekali setahun",
		HostedBy: "Murni Sekali",
		DateStart: time.Time{},
		DateEnd: time.Time{},
		Location: "jakarta",
		Details: "presmian dunia digital",
		Ticket: 5,
	}, nil
}

func (m *mockEvent) Update(id, user_id uint, task *entity.Event) (response.UpdateEvent, error) {
	return response.UpdateEvent{
		Name: task.Name,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (m *mockEvent) Delete(id, user_id uint) (response.DeleteEvent, error) {
	return response.DeleteEvent{
		Name: "event 1",
		DeletedAt: gorm.DeletedAt{},
	}, nil
}

type mockError struct {}

func (m *mockError) Insert(task *entity.Event) response.InsertEvent {
	return response.InsertEvent{}
}

func (m *mockError) GetAll(name, location string, limit, page int) []response.GetEvent {
	return []response.GetEvent{}
}

func (m *mockError) Get(id uint) (response.GetEvent, error) {
	return response.GetEvent{}, errors.New("event not found")
}

func (m *mockError) Update(id, user_id uint, task *entity.Event) (response.UpdateEvent, error) {
	return response.UpdateEvent{}, errors.New("required")
}

func (m *mockError) Delete(id, user_id uint) (response.DeleteEvent, error) {
	return response.DeleteEvent{}, errors.New("you are not allowed to access this resource")
}

type mockErrorForbidden struct {}

func (m *mockErrorForbidden) Insert(task *entity.Event) response.InsertEvent {
	return response.InsertEvent{}
}

func (m *mockErrorForbidden) GetAll(name, location string, limit, page int) []response.GetEvent {
	return []response.GetEvent{}
}

func (m *mockErrorForbidden) Get(id uint) (response.GetEvent, error) {
	return response.GetEvent{}, errors.New("event not found")
}

func (m *mockErrorForbidden) Update(id, user_id uint, task *entity.Event) (response.UpdateEvent, error) {
	return response.UpdateEvent{}, errors.New("you are not allowed to access this resource")
}

func (m *mockErrorForbidden) Delete(id, user_id uint) (response.DeleteEvent, error) {
	return response.DeleteEvent{}, errors.New("you are not allowed to access this resource")
}