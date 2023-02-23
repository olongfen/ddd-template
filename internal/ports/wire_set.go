package ports

import (
	"ddd-template/internal/ports/controller"
	"ddd-template/internal/ports/graph"
	"github.com/google/wire"
)

var Set = wire.NewSet(controller.Set, graph.Set)
