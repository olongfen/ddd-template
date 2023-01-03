//go:build wireinject
// +build wireinject

package main

import (
	"ddd-template/internal/adapters/repository"
	"ddd-template/internal/adapters/store/redis"
	app "ddd-template/internal/application"
	"ddd-template/internal/config"
	"ddd-template/internal/ports/graph"
	"ddd-template/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewServer(configs *config.Configs, logger *zap.Logger) (s *service.Server, fc func()) {
	panic(wire.Build(service.Set, graph.Set, app.Set, redis_store.Set, repository.Set))
}
