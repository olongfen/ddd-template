package rpcx

import "github.com/google/wire"

var Set = wire.NewSet(NewGrpc)
