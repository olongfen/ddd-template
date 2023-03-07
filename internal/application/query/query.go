package query

import "ddd-template/internal/application/query/query_iface"

// Query 查询的数据
type Query struct {
	demo query_iface.IDemoService
}

// Demo 获取demo的查询服务
func (q *Query) Demo() query_iface.IDemoService {
	return q.demo
}

// SetQuery  设置query,
func SetQuery(demo query_iface.IDemoService) *Query {
	return &Query{
		demo: demo,
	}
}
