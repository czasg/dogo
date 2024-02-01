package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"proj/public/httplib"
	"time"
)

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) Create(ctx context.Context, users []User) error {
	return us.DB.WithContext(ctx).CreateInBatches(users, 100).Error
}

func (us *UserService) Delete(ctx context.Context, where map[string]interface{}) error {
	if len(where) < 1 {
		return errors.New("forbid action")
	}
	return us.DB.WithContext(ctx).Where(where).Delete(&User{}).Error
}

func (us *UserService) Update() {}

func (us *UserService) Query(ctx context.Context, query *httplib.QueryParams) ([]User, error) {
	users := []User{}
	err := query.Bind(us.DB.WithContext(ctx)).Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return users, nil
}

type Role struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RoleService struct{}

type UserRole struct {
	ID        int64
	UserID    int64
	RoleID    int64
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
