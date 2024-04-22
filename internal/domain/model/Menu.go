package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"proj/public/httplib"
	"time"
)

// MenuModel
type Menu struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Route     string    `json:"route"`
	Hide      bool      `json:"hide"`
	RootID    int64     `json:"rootID"`
	ParentID  int64     `json:"parentID"`
	Level     int64     `json:"level"`
	OrderID   int64     `json:"orderID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Menu) TableName() string {
	return "menus"
}

func (Menu) New(obj any) (*Menu, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var menu Menu
	err = json.Unmarshal(data, &menu)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

type MenuService struct {
	DB *gorm.DB
}

func (ms *MenuService) Query(ctx context.Context, query *httplib.QueryParams) ([]Menu, error) {
	menu := []Menu{}
	err := query.Bind(ms.DB.WithContext(ctx)).Find(&menu).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return menu, nil
}

func (ms *MenuService) QueryByID(ctx context.Context, id int64) (*Menu, error) {
	menu := Menu{ID: id}
	err := ms.DB.WithContext(ctx).First(&menu).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &menu, nil
}

func (ms *MenuService) Create(ctx context.Context, menu *Menu) error {
	if err := ms.DB.WithContext(ctx).Create(menu).Error; err != nil {
		return fmt.Errorf("failed to create menu: %w", err)
	}
	return nil
}

// RoleMenuModel
type RoleMenu struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	RoleID    int64     `json:"roleID"`
	MenuID    int64     `json:"menuID"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	Menu      Menu      `gorm:"foreignKey:MenuID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (RoleMenu) TableName() string {
	return "role_menus"
}
