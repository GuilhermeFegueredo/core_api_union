package models

import "time"

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// TODO: Status
}

