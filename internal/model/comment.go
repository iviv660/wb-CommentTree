package model

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	ParentID  *int64    `json:"parent_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
