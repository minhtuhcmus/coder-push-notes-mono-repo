package models

import "time"

type User struct {
	ID             int       `json:id`
	Username       string    `json:"username"`
	DisplayName    string    `json:"displayName"`
	PasswordHashed string    `json:"passwordHashed"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
