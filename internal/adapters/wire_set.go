package adapters

import (
	"ddd-template/internal/adapters/repository"
	redis_store "ddd-template/internal/adapters/store/redis"
	"github.com/google/wire"
)

var Set = wire.NewSet(redis_store.Set, repository.Set)
