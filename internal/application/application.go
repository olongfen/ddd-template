package app

import (
	"context"
	"ddd-template/internal/application/mutation"
	"ddd-template/internal/application/query"
)

type Application struct {
	Mutations Mutations
	Queries   Queries
	Cleanup   func()
}

// Mutations 操作变动的数据
type Mutations struct {
	Student mutation.IStudentMutationService
	Class   mutation.IClassMutationService
}

// Queries 查询的数据
type Queries struct {
	Student query.IStudentQueryService
}

func SetQueries(stu query.IStudentQueryService) Queries {
	return Queries{
		Student: stu,
	}
}

func SetMutations(student mutation.IStudentMutationService, class mutation.IClassMutationService) Mutations {
	return Mutations{Student: student, Class: class}
}

func NewApplication(ctx context.Context, mut Mutations, que Queries) Application {
	app := Application{}
	app.Queries = que
	app.Mutations = mut
	app.Cleanup = func() {}
	return app
}
