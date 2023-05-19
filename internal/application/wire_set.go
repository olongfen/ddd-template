package app

import (
	"ddd-template/internal/application/command"
	"ddd-template/internal/application/query"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewApplication, command.Set, query.Set)
