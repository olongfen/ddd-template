package domain

type Field struct {
	Column string // 字段名称
	Value  any    // 字段直
	Symbol string // 符号,<=,<,>,>=,in,like,<>等于默认是""
}

func SetField(col string, val any, symbols ...string) Field {
	symbol := "="
	if len(symbols) > 0 {
		symbol = symbols[0]
	}
	return Field{
		Column: col,
		Value:  val,
		Symbol: symbol,
	}
}

type OtherCond struct {
	PageSize    int
	CurrentPage int
	Sort        []string
	Order       []string
	All         bool
	NoCount     bool
}

type Pagination struct {
	PageSize    int   `json:"pageSize"`
	CurrentPage int   `json:"currentPage"`
	TotalPage   int   `json:"totalPage"`
	TotalCount  int64 `json:"totalCount"`
}
