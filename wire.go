//go:build wireinject
// +build wireinject

package main

import (
	"ddd-template/internal/adapters"
	app "ddd-template/internal/application"
	"ddd-template/internal/rely"
	"ddd-template/internal/service"
	"github.com/google/wire"
)

func NewServer() (s *service.Server, fc func(), err error) {
	panic(wire.Build(service.Set, app.Set, adapters.Set, rely.Set))
}
