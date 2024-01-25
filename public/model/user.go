package model

import "time"

type User struct {
	ID          int
	Name        string
	CreatedTime time.Time
	UpdatedTime time.Time
}

type Role struct {
	ID          int
	Name        string
	CreatedTime time.Time
	UpdatedTime time.Time
}
