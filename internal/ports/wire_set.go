package ports

import (
	"ddd-template/internal/ports/controller"
	"github.com/google/wire"
)

var Set = wire.NewSet(controller.NewHttpServer)
