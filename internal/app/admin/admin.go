package admin

import "github.com/gin-gonic/gin"

func Bind(app gin.IRouter) {
	app.POST("/login")
	app.POST("/logout")
	// user
	app.GET("/users")               // get user list
	app.GET("/user/:uid/details")   // get user details by user-id
	app.POST("/users")              // new user
	app.POST("/user/:uid/details")  // upt user details by user-id
	app.POST("/user/:uid/password") // upt user password by user-id
	app.DELETE("/user/:uid")        // del user by user-id
	// role
	app.GET("/roles")
	app.GET("/user/:uid/roles")
	app.POST("/user/:uid/roles")
	// perm
	app.GET("/permissions")
	app.GET("/role/:rid/permissions")
	app.POST("/role/:rid/permissions")
}
