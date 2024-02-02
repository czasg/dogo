package model

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"proj/public/httplib"
	"time"
)

type User struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Alias     string         `json:"alias"`
	Enable    bool           `json:"enable"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (User) TableName() string {
	return "users"
}

type UserDetail struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	UserID      int64     `json:"userID"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Preference  string    `json:"preference"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	LastLoginAt time.Time `json:"lastLoginAt"`
}

func (UserDetail) TableName() string {
	return "user_details"
}

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) Create(ctx context.Context, user *User, userDetail *UserDetail) error {
	return us.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		userDetail.UserID = user.ID
		if err := tx.Create(userDetail).Error; err != nil {
			return fmt.Errorf("failed to create user detail: %w", err)
		}
		return nil
	})
}

func (us *UserService) DeleteByID(ctx context.Context, ids ...int) error {
	return us.DB.WithContext(ctx).Delete(&User{}, ids).Error
}

func (us *UserService) UpdateUserByID(ctx context.Context, id int64, upt map[string]interface{}) (*User, error) {
	user := User{ID: id}
	if err := us.DB.WithContext(ctx).Clauses(clause.Returning{}).Updates(upt).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

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
