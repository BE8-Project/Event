package response

import "time"

type Comment struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	EventID   uint      `json:"eventId"`
	Field     string    `json:"field"`
	CreatedAt time.Time `json:"created_at"`
}
