package httplib

import (
	"errors"
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

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, HttpResponse{
		Data: data,
	})
}

func Failure(c *gin.Context, err error) {
	var resp HttpResponse
	if errors.As(err, &resp) {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, HttpResponse{
		Code:    99999,
		Message: err.Error(),
	})
}
