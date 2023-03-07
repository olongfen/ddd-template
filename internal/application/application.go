package app

import (
	"ddd-template/internal/application/mutation"
	"ddd-template/internal/application/query"
)

// Application 应用层入口
type Application struct {
	// exec 执行操作入口
	exec *mutation.Mutation
	// query 查询操作入口
	query *query.Query
}

// NewApplication 新建一个应用服务
func NewApplication(mut *mutation.Mutation, que *query.Query) *Application {
	app := &Application{}
	app.query = que
	app.exec = mut
	return app
}

// Exec application exec
func (a *Application) Exec() *mutation.Mutation {
	return a.exec
}

// Query application query
func (a *Application) Query() *query.Query {
	return a.query
}
