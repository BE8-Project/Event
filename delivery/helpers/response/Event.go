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
	Name       string `json:"name" validate:"required"`
	HostedBy   string `json:"hosted_by" validate:"required"`
	DateStart  time.Time `json:"date_start" validate:"required"`
	DateEnd    time.Time `json:"date_end" validate:"required"`
	Location   string `json:"location" validate:"required"`
	Details    string `json:"details" validate:"required"`
	Ticket     int    `json:"ticket" validate:"required"`
}

type UpdateEvent struct {
	Name      string `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteEvent struct {
	Name      string `json:"name"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}