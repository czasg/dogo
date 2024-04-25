package service

import (
	"context"
	"proj/internal/domain/model"
)

type Menus struct {
	model.Menu
	Submenu []Menus `json:"submenu"`
}

type MenuService struct {
	UserService     model.UserService
	RoleService     model.RoleService
	MenuService     model.MenuService
	RoleMenuService model.RoleMenuService
}

func (ms *MenuService) Menus(menus []model.Menu) []Menus {
	var ans []Menus
	for _, menu := range menus {
		if menu.Level == 0 {
			ans = append(ans, Menus{Menu: menu})
			continue
		}
		if menu.Level == 1 {
			for ii, rootMenu := range ans {
				if rootMenu.ID == menu.ParentID {
					rootMenu.Submenu = append(rootMenu.Submenu, Menus{Menu: menu})
					ans[ii] = rootMenu
					break
				}
			}
			continue
		}
	}
	return ans
}

func (ms *MenuService) MenusByUserID(ctx context.Context, id int64) ([]Menus, error) {
	roles, err := ms.UserService.QueryUserRoleByID(ctx, id)
	if err != nil {
		return nil, nil
	}
	rid := []int64{}
	for _, role := range roles {
		rid = append(rid, role.ID)
	}
	menus, err := ms.RoleMenuService.GetMenusByRoleID(ctx, rid...)
	if err != nil {
		return nil, nil
	}
	return ms.Menus(menus), nil
}
