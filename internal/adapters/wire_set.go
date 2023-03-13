package adapters

import (
	"ddd-template/internal/adapters/repository"
	"ddd-template/internal/adapters/store"
	"github.com/google/wire"
)

var Set = wire.NewSet(store.Set, repository.Set)
