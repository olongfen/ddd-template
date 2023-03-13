package store

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewRedisStore)
