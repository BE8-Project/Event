package response

import "time"

type InsertEvent struct {
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}