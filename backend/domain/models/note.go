package models

import "time"

type Note struct {
	ID        int       `json:"ID"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Active    bool      `json:"active"`
	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
