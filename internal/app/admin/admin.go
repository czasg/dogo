package admin

import (
	"github.com/gin-gonic/gin"
	v1 "proj/internal/app/admin/v1"
)

func Bind(app gin.IRouter) {
	// admin group v1
	admin := app.Group("/app-admin/v1")
	{
		admin.POST("/login", v1.Login)   // user login
		admin.POST("/logout", v1.Logout) // user logout
	}
	{
		// user
		admin.GET("/users", v1.UserList)  // get user list
		admin.GET("/user/:uid/details")   // get user details by user-id
		admin.POST("/users")              // new user
		admin.POST("/user/:uid/details")  // upt user details by user-id
		admin.POST("/user/:uid/password") // upt user password by user-id
		admin.POST("/user/:uid/role")     // upt user role by user-id
		admin.POST("/user/:uid/enable")   // upt user enable by user-id
		admin.DELETE("/user/:uid")        // del user by user-id
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
