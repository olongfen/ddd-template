package service

import (
	"ddd-template/internal/ports/controller"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewServer, controller.Set)
