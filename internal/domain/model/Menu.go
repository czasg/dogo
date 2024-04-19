package model

import "time"

// MenuModel
type Menu struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Route     string    `json:"route"`
	Hide      bool      `json:"hide"`
	Level     int64     `json:"level"`
	Order     int64     `json:"order"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Menu) TableName() string {
	return "menus"
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
