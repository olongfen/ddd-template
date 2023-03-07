package app

import (
	"ddd-template/internal/application/mutation/mutat_iface"
	"ddd-template/internal/application/query/query_iface"
)

// Application 应用层入口
type Application struct {
	// exec 执行操作入口
	exec *Mutation
	// query 查询操作入口
	query *Query
}

// Mutation 操作变动的数据
type Mutation struct {
	demo mutat_iface.IDemoService
}

// Demo 获取demo的写入服务
func (m *Mutation) Demo() mutat_iface.IDemoService {
	return m.demo
}

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

// SetMutation set mutation
func SetMutation(demo mutat_iface.IDemoService) *Mutation {
	return &Mutation{demo: demo}
}

// NewApplication 新建一个应用服务
func NewApplication(mut *Mutation, que *Query) *Application {
	app := &Application{}
	app.query = que
	app.exec = mut
	return app
}

// Exec application exec
func (a *Application) Exec() *Mutation {
	return a.exec
}

// Query application query
func (a *Application) Query() *Query {
	return a.query
}
