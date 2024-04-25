package admin

import (
	"github.com/gin-gonic/gin"
	v1 "proj/internal/app/admin/v1"
)

func Bind(app gin.IRouter) {
	// admin group v1
	admin := app.Group("/app-admin/v1")
	{
		ua := v1.DefaultUserApp()
		// user
		admin.POST("/login", ua.Login)                           // user login
		admin.POST("/logout", ua.Logout)                         // user logout
		admin.GET("/users", ua.UserList)                         // get user list
		admin.GET("/user/:uid/details", ua.UserDetails)          // get user details by user-id
		admin.POST("/users", ua.CreateUser)                      // new user
		admin.POST("/user/:uid/details", ua.UpdateUserDetail)    // upt user details by user-id
		admin.POST("/user/:uid/password", ua.UpdateUserPassword) // upt user password by user-id
		admin.GET("/user/:uid/roles", ua.UserRoleList)           // get user role by user-id
		admin.POST("/user/:uid/roles", ua.UpdateUserRole)        // upt user role by user-id
		admin.GET("/user/:uid/menus", ua.GetUserMenu)            // get user menu by user-id
		admin.POST("/user/:uid/enable", ua.UpdateUserEnable)     // upt user enable by user-id
		admin.DELETE("/user/:uid/record", ua.DeleteUser)         // del user by user-id
	}
	{
		ra := v1.DefaultRoleApp()
		// role
		admin.GET("/roles", ra.List)                           // get role list
		admin.POST("/roles", ra.Create)                        // mew role
		admin.POST("/role/:rid/details", ra.UpdateRoleDetails) // upt role details by role-id
		admin.GET("/role/:rid/menus", ra.GetRoleMenus)         // get role menu by role-id
		admin.POST("/role/:rid/menus", ra.UpdateMenus)         // upt role menu by role-id
		admin.DELETE("/role/:rid/record", ra.DeleteRole)       // del role by role-id
	}
	{
		ma := v1.DefaultMenuApp()
		// menu
		admin.GET("/menus", ma.MenuList)                           // get menu list
		admin.POST("/menus", ma.CreateMenu)                        // new root menu
		admin.POST("/menu/:mid/secondary", ma.CreateSecondaryMenu) // new secondary menu by menu-id
		admin.GET("/menu/:mid/details", ma.MenuDetails)            // upt menu details by menu-id
	}
}
