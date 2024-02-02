package model

import (
	"context"
	"errors"
	"fmt"
	"proj/public/httplib"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := us.DB.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	userDetail.UserID = user.ID
	if err := us.DB.WithContext(ctx).Create(userDetail).Error; err != nil {
		return fmt.Errorf("failed to create user detail: %w", err)
	}
	return nil
}

func (us *UserService) DeleteByID(ctx context.Context, ids ...int) error {
	return us.DB.WithContext(ctx).Delete(&User{}, ids).Error
}

func (us *UserService) UpdateUserByID(ctx context.Context, id int64, upt map[string]interface{}) (*User, error) {
	if len(upt) < 1 {
		return nil, errors.New("invalid action")
	}
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

func (us *UserService) QueryByID(ctx context.Context, id int64) (*User, *UserDetail, error) {
	user := User{ID: id}
	userDetail := UserDetail{UserID: id}
	if err := us.DB.WithContext(ctx).First(&user).Error; err != nil {
		return nil, nil, err
	}
	if err := us.DB.WithContext(ctx).First(&userDetail).Error; err != nil {
		return nil, nil, err
	}
	return &user, &userDetail, nil
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
