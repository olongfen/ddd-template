package domain

// Field 数据库字段
type Field struct {
	Column string // 字段名称
	Value  any    // 字段直
	Symbol string // 符号,<=,<,>,>=,in,like,<>等于默认是""
}

// SetField 设置数据库字段 样式 where col symbol val
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

// OtherCond 其他条件定义
type OtherCond struct {
	// PageSize limit
	PageSize int
	// CurrentPage 当前页
	CurrentPage int
	// Sort 排序字段
	Sort []string
	// Order 排序值 desc和asc,长度跟sort一样，index一一对应
	Order []string
	// All true获取全部数据
	All bool
	// NoCount 不需要执行count 查询
	NoCount bool
}

// Pagination 页码数据
type Pagination struct {
	PageSize    int   `json:"pageSize"`
	CurrentPage int   `json:"currentPage"`
	TotalPage   int   `json:"totalPage"`
	TotalCount  int64 `json:"totalCount"`
}
