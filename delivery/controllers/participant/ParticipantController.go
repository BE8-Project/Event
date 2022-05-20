package participant

import (
	"event/config"
	"event/delivery/helpers/request"
	"event/delivery/helpers/response"
	"event/delivery/middlewares"
	"event/delivery/usecase"
	"event/entity"
	repoParticipant "event/repository/participant"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type participantController struct {
	Connect  repoParticipant.ParticipantModel
	Validate *validator.Validate
}

func NewParticipantController(conn repoParticipant.ParticipantModel, valid *validator.Validate) *participantController {
	return &participantController{
		Connect:  conn,
		Validate: valid,
	}
}

func (c *participantController) Insert() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(ctx))
		var request request.InsertParticipant

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestBind(err))
		}

		if err := c.Validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		order := entity.Participant{
			EventID:		request.EventID,
			PaymentType:    request.PaymentType,
			Total:          request.Total,
			OrderID:		"DM-" + usecase.Random(),
			Status:         "pending",
			UserID:         user_id,
		}

		result, err := c.Connect.Insert(&order)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		config.SetupGlobalMidtransConfigApi()
		midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")
		midtrans.SetPaymentOverrideNotification("https://midtrans-java.herokuapp.com/notif/override")

		resp, _ := coreapi.ChargeTransactionWithMap(usecase.Gopay(result.OrderID, result.Total))

		var message []interface{}
		var transaction_id string
		for key, value := range resp {
			if key == "actions" {
				message = value.([]interface{})
			}

			if key == "transaction_id" {
				transaction_id = value.(string)
			}
		}

		var action map[string]interface{}
		for key, value := range message {
			if key == 1 {
				action = value.(map[string]interface{})
			}
		}

		var data map[string]interface{} = make(map[string]interface{})
		data["order_id"] = result.OrderID
		data["payment_type"] = "gopay"
		data["total"] = result.Total
		data["status"] = "pending"
		data["payment_simulator"] = action["url"]
		data["payment_url"] = "https://api.sandbox.midtrans.com/v2/gopay/" + transaction_id + "/qr-code"
		data["created_at"] = result.CreatedAt

		return ctx.JSON(http.StatusCreated, response.StatusCreated("success create Order!", data))
	}
}

func (c *participantController) GetStatus() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(ctx))
		order_id := ctx.Param("order_id")

		config.SetupGlobalMidtransConfigApi()
		midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")
		midtrans.SetPaymentOverrideNotification("https://midtrans-java.herokuapp.com/notif/override")

		resp, _ := coreapi.CheckTransaction(order_id)

		update := entity.Participant{
			OrderID:	resp.OrderID,
			Status:		resp.TransactionStatus,
		}

		result, err := c.Connect.Update(order_id, user_id, &update)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("success get Status!", result))
	}
}

func (c *participantController) Cancel() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(ctx))
		order_id := ctx.Param("order_id")

		config.SetupGlobalMidtransConfigApi()
		midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")
		midtrans.SetPaymentOverrideNotification("https://midtrans-java.herokuapp.com/notif/override")

		resp, _ := coreapi.CancelTransaction(order_id)

		update := entity.Participant{
			OrderID:	resp.OrderID,
			Status:		resp.TransactionStatus,
		}

		result, err := c.Connect.Update(order_id, user_id, &update)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("success cancel Order!", result))
	}
}

func (c *participantController) GetByUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := uint(middlewares.ExtractTokenUserId(ctx))

		results := c.Connect.GetByUser(user_id)

		if len(results) == 0 {
			return ctx.JSON(http.StatusNotFound, response.StatusNotFound("Data tidak ditemukan!"))
		}

		return ctx.JSON(http.StatusOK, response.StatusOK("Berhasil mengambil myEvent!", results))
	}
}