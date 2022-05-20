package response

import (
	"time"
)

type InsertParticipant struct {
	OrderID string `json:"order_id"`
	Total int `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateParticipat struct {
	OrderID string `json:"order_id"`
	Status string `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetParticipant struct {
	ID		   uint `json:"id"`
	EventID     uint `json:"event_id"`
	OrderID     string `json:"order_id"`
	PaymentType string `json:"payment_type"`
	Total       int `json:"total"`
	Status      string `json:"status"`
	Event 		[]GetEvent `json:"event"`
}