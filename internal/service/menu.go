package service

import (
	"proj/internal/domain/model"
)

type MenuService struct{}

type Menus struct {
	model.Menu
	Submenu []Menus `json:"submenu"`
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
