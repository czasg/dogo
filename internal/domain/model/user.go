package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) List() ([]User, error) {
	return nil, nil
}

func (us *UserService) Create() {}

func (us *UserService) Update() {}

func (us *UserService) Delete() {}
