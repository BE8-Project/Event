package request

type InsertEvent struct {
	Name       string `json:"name" validate:"required"`
	HostedBy   string `json:"hosted_by" validate:"required"`
	DateStart  string `json:"date_start" validate:"required"`
	DateEnd    string `json:"date_end" validate:"required"`
	Location   string `json:"location" validate:"required"`
	Details    string `json:"details" validate:"required"`
	Ticket     int    `json:"ticket" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}
