package controller

import (
	"ddd-template/internal/service/controller/handler"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewHTTPServer, handler.Set, NewMiddleware)
