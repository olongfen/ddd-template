//go:build wireinject
// +build wireinject

package main

import (
	"ddd-template/internal/adapters"
	app "ddd-template/internal/application"
	"ddd-template/internal/ports"
	"ddd-template/internal/rely"
	"ddd-template/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewServer(configs *rely.Configs, logger *zap.Logger) (s *service.Server, fc func()) {
	panic(wire.Build(service.Set, ports.Set, app.Set, adapters.Set))
}
