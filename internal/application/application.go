package app

import (
	"context"
	"ddd-template/internal/application/mutation"
)

type Application struct {
	Mutations Mutations
	Queries   Queries
	Cleanup   func()
}

// Mutations 操作变动的数据
type Mutations struct {
	Student mutation.IStudentMutationService
}

// Queries 查询的数据
type Queries struct {
}

func SetQueries() Queries {
	return Queries{}
}

func SetMutations(student mutation.IStudentMutationService) Mutations {
	return Mutations{Student: student}
}

func NewApplication(ctx context.Context, mut Mutations, que Queries) Application {
	app := Application{}
	app.Queries = que
	app.Mutations = mut
	app.Cleanup = func() {}
	return app
}
