package request

type InsertEvent struct {
	Name       string    `json:"name" validate:"required"`
	HostedBy   string    `json:"hosted_by" validate:"required"`
	Date       string `json:"date" validate:"required"`
	Location   string    `json:"location" validate:"required"`
	Details    string    `json:"details" validate:"required"`
	Ticket     int       `json:"ticket" validate:"required"`
	CategoryID uint      `json:"category_id" validate:"required"`
}