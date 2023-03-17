package graph

import "github.com/google/wire"

var Set = wire.NewSet(NewResolver)
