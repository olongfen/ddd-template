package ports

import "github.com/google/wire"

var Set = wire.NewSet(NewHttpServer)
