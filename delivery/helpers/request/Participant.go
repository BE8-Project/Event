package request

type InsertParticipant struct {
	EventID     uint   `json:"event_id" validate:"required"`
	PaymentType string `json:"payment_type" validate:"required"`
	Total       int    `json:"total" validate:"required"`
}