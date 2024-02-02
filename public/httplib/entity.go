package httplib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Err     error       `json:"-"`
}

func (hr HttpResponse) Error() string {
	if hr.Err != nil {
		return hr.Err.Error()
	}
	return ""
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, HttpResponse{
		Data: data,
	})
}

func ERROR(c *gin.Context, err error) {
}
