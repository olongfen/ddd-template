package app

import (
	"ddd-template/internal/application/mutation"
	"ddd-template/internal/application/query"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewApplication, SetMutations, SetQueries, mutation.NewUserMutation, mutation.NewClassMutation, query.NewQueryStudent, query.NewQueryClass)
