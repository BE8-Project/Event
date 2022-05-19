package response

import "time"

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