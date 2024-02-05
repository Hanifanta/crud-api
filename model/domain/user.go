package domain

import "time"

type User struct {
	Id        string
	Email     string
	Username  string
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
