package httplib

import "gorm.io/gorm"

type QueryParams struct {
	PageSize int                    `json:"pageSize"`
	PageNum  int                    `json:"pageNum"`
	Sort     string                 `json:"sort"`
	Where    map[string]interface{} `json:"-"`
}

func (q *QueryParams) Bind(db *gorm.DB) *gorm.DB {
	return db.Where(q.Where).Order(q.Sort).Limit(min(max(q.PageSize, 10), 100)).Offset(q.PageNum)
}
