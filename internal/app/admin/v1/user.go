package v1

import (
	"net/http"
	"proj/internal/domain/model"
	"proj/public/httplib"

	"github.com/gin-gonic/gin"
)

type UserApp struct {
	userService model.UserService
}

func (ua *UserApp) UserList(c *gin.Context) {
	query := httplib.QueryParams{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, httplib.HttpResponse{})
		return
	}
	users, err := ua.userService.Query(c, &query)
	if err != nil {
		c.JSON(http.StatusOK, httplib.HttpResponse{})
		return
	}
	c.JSON(http.StatusOK, httplib.HttpResponse{Data: users})
}

func (ua *UserApp) UserDetails(c *gin.Context) {
	query := httplib.QueryParams{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, httplib.HttpResponse{})
		return
	}
	users, err := ua.userService.Query(c, &query)
	if err != nil {
		c.JSON(http.StatusOK, httplib.HttpResponse{})
		return
	}
	c.JSON(http.StatusOK, httplib.HttpResponse{Data: users})
}
