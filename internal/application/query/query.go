package query

// Query 查询的数据
type Query struct {
	demo IDemoService
}

// Demo 获取demo的查询服务
func (q *Query) Demo() IDemoService {
	return q.demo
}

// SetQuery  设置query,
func SetQuery(demo IDemoService) *Query {
	return &Query{
		demo: demo,
	}
}
