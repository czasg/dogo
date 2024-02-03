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
		admin.POST("/user/:uid/name", ua.UpdateUserDetail)       // upt user details by user-id
		admin.POST("/user/:uid/details", ua.UpdateUserDetail)    // upt user details by user-id
		admin.POST("/user/:uid/password", ua.UpdateUserPassword) // upt user password by user-id
		admin.POST("/user/:uid/role", ua.UpdateUserRole)         // upt user role by user-id
		admin.POST("/user/:uid/enable", ua.UpdateUserEnable)     // upt user enable by user-id
		admin.DELETE("/user/:uid/record", ua.DeleteUser)         // del user by user-id
	}
	{
		// role
		admin.GET("/roles")
		admin.GET("/user/:uid/roles")
		admin.POST("/user/:uid/roles")
	}
	{
		// perm
		admin.GET("/permissions")
		admin.GET("/role/:rid/permissions")
		admin.POST("/role/:rid/permissions")

	}
	{
		// menu
		admin.GET("/menus") // get menu list
	}
	{
		// apis
		admin.GET("/apis") // get api list
	}
}
