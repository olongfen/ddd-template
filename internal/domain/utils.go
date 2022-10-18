package domain

type Field struct {
	Column string // 字段名称
	Value  any    // 字段直
	Symbol string // 符号,<=,<,>,>=,in,like,<>等于默认是""
}

type OtherCond struct {
	PageSize    int
	CurrentPage int
	Sort        []string
	Order       []string
}

type Pagination struct {
	PageSize    int   `json:"pageSize"`
	CurrentPage int   `json:"currentPage"`
	TotalPage   int   `json:"totalPage"`
	TotalCount  int64 `json:"totalCount"`
}
