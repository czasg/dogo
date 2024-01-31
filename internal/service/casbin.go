package service

import "proj/internal/domain/model"

type CasbinService struct {
	model.UserService
	model.RoleService
}
