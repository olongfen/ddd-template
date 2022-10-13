//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"ddd-template/internal/adapters/handler"
	"ddd-template/internal/adapters/repository"
	app "ddd-template/internal/application"
	"ddd-template/internal/config"
	"ddd-template/internal/domain"
	"ddd-template/internal/ports"
	"ddd-template/internal/ports/controller"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewServer(ctx context.Context, configs *config.Configs, logger *zap.Logger) (s controller.HttpServer, fc func()) {
	panic(wire.Build(ports.Set, handler.Set, app.Set, domain.Set, repository.Set))
}
