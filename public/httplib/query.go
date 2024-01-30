package httplib

import "github.com/gin-gonic/gin"

type QueryParams struct {
	PageSize int    `json:"pageSize"`
	PageNum  int    `json:"pageNum"`
	Sort     string `json:"sort"`
}

func Query(c *gin.Context) {}
