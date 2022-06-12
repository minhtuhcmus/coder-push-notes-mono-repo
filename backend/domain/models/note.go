package models

import "time"

type Note struct {
	ID        int       `json:"ID"`
	Text      string    `json:"text"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
