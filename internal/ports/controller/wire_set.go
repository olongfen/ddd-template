package controller

import (
	"ddd-template/internal/ports/controller/handler"
	"ddd-template/internal/ports/controller/middleware"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewHttpServer, handler.Set, middleware.NewMiddleware)
