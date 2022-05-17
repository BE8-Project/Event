package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(35);not null"`
	Username string  `gorm:"type:varchar(35);not null;unique"`
	Email    string  `gorm:"type:varchar(100);not null;unique"`
	HP       string  `gorm:"type:varchar(20);not null;unique"`
	Password string  `gorm:"type:varchar(255);not null"`
	Role     int     `gorm:"type:int;not null"`
	Events   []Event `gorm:"foreignkey:Username;references:Username"`
}

type Category struct {
	gorm.Model
	Name   string
	Events []Event `gorm:"foreignkey:CategoryID"`
}

type Event struct {
	gorm.Model
	Name       string
	HostedBy   string
	Date       time.Time
	Location   string
	Details    string
	Ticket     int
	Username   string
	CategoryID uint
}
