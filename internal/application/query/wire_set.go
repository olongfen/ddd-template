package query

import "github.com/google/wire"

var Set = wire.NewSet(SetQuery, NewDemo)
