package model

import "time"

type Organization struct {
	ID                 int64
	Name               string
	Alias              string
	RootOrganizationID int64
	PreOrganizationID  int64
	Level              int64
	Description        string
	CreatedAt          time.Time    `json:"createdAt"`
	UpdatedAt          time.Time    `json:"updatedAt"`
	Departments        []Department `gorm:"many2many:organization_departments;"`
}

type Department struct {
	ID            int64
	Name          string
	Alias         string
	Description   string
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	Organizations []Organization `gorm:"many2many:organization_departments;"`
}

type OrganizationDepartment struct {
	ID             int64
	OrganizationID int64
	DepartmentID   int64
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
