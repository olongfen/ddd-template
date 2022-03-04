//go:build wireinject
// +build wireinject

package main

import (
	"ddd-template/adapters/repositry"
	"ddd-template/adapters/restful"
	"ddd-template/app"
	"ddd-template/app/serve"
	"ddd-template/common/conf"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func Init(cfg conf.Configs, log *zap.Logger) *app.Application {
	wire.Build(app.ProviderSet, restful.ProviderSet, serve.ProviderSet, repositry.ProviderSet)
	return nil
}
