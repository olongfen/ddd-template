//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"ddd-template/internal/adapters/respository"
	app "ddd-template/internal/application"
	"ddd-template/internal/config"
	"ddd-template/internal/domain"
	"ddd-template/internal/ports"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewServer(ctx context.Context, configs *config.Configs, logger *zap.Logger) (s ports.HttpServer, fc func()) {
	panic(wire.Build(ports.Set, app.Set, domain.Set, respository.Set))
}
