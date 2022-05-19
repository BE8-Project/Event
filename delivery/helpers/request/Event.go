package request

type InsertEvent struct {
	Name       string `json:"name" validate:"required" form:"name"`
	HostedBy   string `json:"hosted_by" validate:"required" form:"hosted_by"`
	DateStart  string `json:"date_start" validate:"required" form:"date_start"`
	DateEnd    string `json:"date_end" validate:"required" form:"date_end"`
	Location   string `json:"location" validate:"required" form:"location"`
	Details    string `json:"details" validate:"required" form:"details"`
	Ticket     int    `json:"ticket" validate:"required" form:"ticket"`
	CategoryID uint   `json:"category_id" validate:"required" form:"category_id"`
}

type UpdateEvent struct {
	Name       string `json:"name"`
	HostedBy   string `json:"hosted_by"`
	DateStart  string `json:"date_start"`
	DateEnd    string `json:"date_end"`
	Location   string `json:"location"`
	Details    string `json:"details"`
	Ticket     int    `json:"ticket"`
	CategoryID uint   `json:"category_id"`
}