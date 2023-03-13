package schema

import (
	"github.com/olongfen/toolkit/multi/xerror"
)

type QueryOptions struct {
	CurrentPage int `json:"currentPage" validate:"min=0" query:"currentPage"`
	PageSize    int `json:"pageSize" validate:"min=1,max=100" query:"pageSize"`
	// sort 忽略下面两个字段自动生成文档
	Sort []string `json:"-" query:"sort"`
	// order
	Order []string `json:"-" query:"order"`
}

func (q QueryOptions) Validate(language string) (err error) {
	if len(q.Sort) != len(q.Order) {
		err = xerror.NewError(xerror.SortParameterMismatch, language)
		return
	}
	return
}

type Pagination struct {
	PageSize    int   `json:"pageSize"`
	CurrentPage int   `json:"currentPage"`
	TotalPage   int   `json:"totalPage"`
	TotalCount  int64 `json:"totalCount"`
}
