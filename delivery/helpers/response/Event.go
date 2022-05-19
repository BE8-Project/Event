package response

import (
	"time"

	"gorm.io/gorm"
)

type InsertEvent struct {
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetEvent struct {
	ID		   uint `json:"id"`
	Name       string `json:"name"`
	HostedBy   string `json:"hosted_by"`
	DateStart  time.Time `json:"date_start"`
	DateEnd    time.Time `json:"date_end"`
	Location   string `json:"location"`
	Details    string `json:"details"`
	Ticket     int    `json:"ticket"`
}

type UpdateEvent struct {
	Name      string `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteEvent struct {
	Name      string `json:"name"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}