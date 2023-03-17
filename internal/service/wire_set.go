package service

import (
	"ddd-template/internal/service/controller"
	"ddd-template/internal/service/graph"
	"github.com/google/wire"
)

var Set = wire.NewSet(controller.Set, graph.Set, NewServer)
