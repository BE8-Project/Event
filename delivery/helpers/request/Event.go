package request

import "time"

type InsertEvent struct {
	Name       string    `json:"name"`
	HostedBy   string    `json:"hosted_by"`
	Date       time.Time `json:"date"`
	Location   string    `json:"location"`
	Details    string    `json:"details"`
	Ticket     int       `json:"ticket"`
	CategoryID uint      `json:"category_id"`
}