package entities

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUser struct {
	Name string `json:"name"`
}
