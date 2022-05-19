package request

type InsertEvent struct {
	Name       string `json:"name" validate:"required" form:"name"`
	HostedBy   string `json:"hosted_by" validate:"required" form:"hosted_by"`
	DateStart  string `json:"date_start" validate:"required" form:"date_start"`
	DateEnd    string `json:"date_end" validate:"required" form:"date_end"`
	Location   string `json:"location" validate:"required" form:"location"`
	Details    string `json:"details" validate:"required" form:"details"`
	Ticket     int    `json:"ticket" validate:"required" form:"ticket"`
	Price      int    `json:"price" validate:"required" form:"price"`
	CategoryID uint   `json:"category_id" validate:"required" form:"category_id"`
}

type UpdateEvent struct {
	Name       string `json:"name" form:"name"`
	HostedBy   string `json:"hosted_by" form:"hosted_by"`
	DateStart  string `json:"date_start" form:"date_start"`
	DateEnd    string `json:"date_end" form:"date_end"`
	Location   string `json:"location" form:"location"`
	Details    string `json:"details" form:"details"`
	Ticket     int    `json:"ticket" form:"ticket"`
	Price      int    `json:"price" form:"price"`
	CategoryID uint   `json:"category_id" form:"category_id"`
}