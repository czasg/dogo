package model

import (
	"context"
	"errors"
	"fmt"
	"proj/public/httplib"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	Enable    bool      `json:"enable"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

func (us *UserService) DeleteByID(ctx context.Context, ids ...int64) error {
	return us.DB.WithContext(ctx).Delete(&User{}, ids).Error
}

func (us *UserService) UpdateUserByID(ctx context.Context, id int64, upt map[string]interface{}) (*User, error) {
	if len(upt) < 1 {
		return nil, errors.New("invalid action")
	}
	user := User{ID: id}
	if err := us.DB.WithContext(ctx).Model(&user).Updates(upt).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) UpdateUserDetailByUserID(ctx context.Context, id int64, upt map[string]interface{}) (*UserDetail, error) {
	if len(upt) < 1 {
		return nil, errors.New("invalid action")
	}
	userDetail := UserDetail{UserID: id}
	err := us.DB.WithContext(ctx).Model(&userDetail).Where("user_id = ?", id).Updates(upt).Error
	if err != nil {
		return nil, err
	}
	return &userDetail, nil
}

func (us *UserService) UpdateUserRoleByID(ctx context.Context, uid int64, rid []int64) error {
	if err := us.DB.WithContext(ctx).Where("user_id = ?", uid).Delete(&UserRole{}).Error; err != nil {
		return err
	}
	ur := []UserRole{}
	for _, id := range rid {
		ur = append(ur, UserRole{
			UserID: uid,
			RoleID: id,
		})
	}
	if err := us.DB.WithContext(ctx).CreateInBatches(ur, 100).Error; err != nil {
		return err
	}
	return nil
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
	if err := us.DB.WithContext(ctx).Where("user_id = ?", id).First(&userDetail).Error; err != nil {
		return nil, nil, err
	}
	return &user, &userDetail, nil
}

func (us *UserService) QueryByName(ctx context.Context, name string) (*User, *UserDetail, error) {
	user := User{}
	userDetail := UserDetail{}
	err := us.DB.WithContext(ctx).Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, fmt.Errorf("user[%s] not found", name)
		}
		return nil, nil, err
	}
	if err := us.DB.WithContext(ctx).Where("user_id = ?", user.ID).First(&userDetail).Error; err != nil {
		return nil, nil, err
	}
	return &user, &userDetail, nil
}

func (us *UserService) ExistsByUserID(ctx context.Context, ids ...int64) (bool, error) {
	users := []User{}
	if err := us.DB.WithContext(ctx).Select([]string{"id"}).Find(users, ids).Error; err != nil {
		return false, err
	}
	return len(users) == len(ids), nil
}

type Role struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RoleService struct {
	DB *gorm.DB
}

func (rs *RoleService) Query(ctx context.Context, query *httplib.QueryParams) ([]Role, error) {
	roles := []Role{}
	err := query.Bind(rs.DB.WithContext(ctx)).Find(&roles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return roles, nil
}

func (rs *RoleService) ExistsByRoleID(ctx context.Context, ids ...int64) (bool, error) {
	roles := []Role{}
	if err := rs.DB.WithContext(ctx).Select([]string{"id"}).Find(roles, ids).Error; err != nil {
		return false, err
	}
	return len(roles) == len(ids), nil
}

/*
CREATE TABLE user_roles (
    user_id INT,
    role_id INT,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (role_id) REFERENCES role(id)
);
*/

type UserRole struct {
	UserID    int64     `gorm:"primaryKey;autoIncrement:false"`
	RoleID    int64     `gorm:"primaryKey;autoIncrement:false"`
	User      User      `gorm:"foreignKey:UserID"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
