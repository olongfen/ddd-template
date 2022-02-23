//go:build wireinject
// +build wireinject

package main

import (
	"ddd-template/adapters/repositry"
	"ddd-template/adapters/restful"
	"ddd-template/application"
	"ddd-template/domain/service"
	"ddd-template/infra/conf"
	"ddd-template/infra/database"
	"ddd-template/infra/xgin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func Init(cfg conf.Configs, log *zap.Logger) *restful.Rest {
	wire.Build(restful.ProviderSet, xgin.ProviderSet, application.ProviderSet, service.ProviderSet, repositry.ProviderSet, database.ProviderSet)
	return nil
}
