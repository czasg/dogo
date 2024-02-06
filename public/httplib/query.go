package httplib

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type QueryMapping struct {
	StringMap QueryStringMap
	IntMap    QueryIntMap
}

func (q QueryMapping) Parse(c *gin.Context) (*QueryParams, error) {
	params := QueryParams{Where: map[string]interface{}{}}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		return nil, err
	}
	err = q.StringMap.Parse(c, params.Where)
	if err != nil {
		return nil, err
	}
	err = q.IntMap.Parse(c, params.Where)
	if err != nil {
		return nil, err
	}
	return &params, nil
}

type QueryStringMap struct {
	Eq map[string]string
	Lk map[string]string
}

type QueryIntMap struct {
	Eq map[string]string
	Gt map[string]string
	Ge map[string]string
	Lt map[string]string
	Le map[string]string
}

func (q QueryIntMap) Parse(c *gin.Context, where map[string]interface{}) error {
	for symbol, opMap := range map[string]map[string]string{
		"=":  q.Eq,
		">":  q.Gt,
		">=": q.Ge,
		"<":  q.Lt,
		"<=": q.Le,
	} {
		for k, sub := range opMap {
			v, ok := c.GetQuery(k)
			if !ok {
				continue
			}
			v = strings.TrimSpace(v)
			if strings.Contains(v, "-") || strings.Contains(v, "=") {
				return errors.New("illegal string")
			}
			i, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				return errors.New("illegal string")
			}
			if strings.Contains(sub, "?") {
				where[sub] = i
			} else {
				where[fmt.Sprintf("%s %s ?", sub, symbol)] = i
			}
		}
	}
	return nil
}

func (q QueryStringMap) Parse(c *gin.Context, where map[string]interface{}) error {
	for k, sub := range q.Eq {
		v, ok := c.GetQuery(k)
		if !ok {
			continue
		}
		v = strings.TrimSpace(v)
		if strings.Contains(v, "-") {
			return errors.New("illegal string")
		}
		if strings.Contains(sub, "?") || strings.Contains(v, "=") {
			where[sub] = v
		} else {
			where[fmt.Sprintf("%s = ?", sub)] = v
		}
	}
	for k, sub := range q.Lk {
		v, ok := c.GetQuery(k)
		if !ok {
			continue
		}
		v = strings.TrimSpace(v)
		if strings.Contains(v, "-") || strings.Contains(v, "=") {
			return errors.New("illegal string")
		}
		v = "%" + v + "%"
		if strings.Contains(sub, "?") {
			where[sub] = v
		} else {
			where[fmt.Sprintf("%s like ?", sub)] = v
		}
	}
	return nil
}

type QueryParams struct {
	PageSize int                    `json:"pageSize" form:"pageSize"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	Sort     string                 `json:"sort" form:"sort"`
	Where    map[string]interface{} `json:"-" form:"-"`
}

func (q *QueryParams) Bind(db *gorm.DB) *gorm.DB {
	return db.Where(q.Where).Order(q.Sort).Limit(min(max(q.PageSize, 10), 100)).Offset(q.PageNum)
}
