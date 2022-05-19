package usecase

import (
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

func Random() string {
	time.Sleep(500 * time.Millisecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func Gopay(order_id string, total int) *coreapi.ChargeReqWithMap {
	req := &coreapi.ChargeReqWithMap{
		"payment_type": "gopay",
		"transaction_details": map[string]interface{}{
			"order_id":     order_id,
			"gross_amount": total,
		},
	}

	return req
}