package xgin

import "github.com/google/wire"

var Set = wire.NewSet(NewHTTPServer)
