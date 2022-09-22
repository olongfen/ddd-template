package app

import (
	"ddd-template/internal/application/mutation"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewApplication, SetMutations, SetQueries, mutation.NewUserMutation)
