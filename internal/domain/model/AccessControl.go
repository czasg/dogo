package model

import (
	"time"
)

type AccessControl struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Typ       string    `json:"typ"` // u:user r:role o:organization d:department
	V1        string    `json:"v1"`
	V2        string    `json:"v2"`
	V3        string    `json:"v3"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Api struct{}

type Menu struct {
}
