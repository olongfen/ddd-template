// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-template/adapters/repositry"
	"ddd-template/adapters/restful"
	"ddd-template/application"
	"ddd-template/domain/service"
	"ddd-template/infra/conf"
	"ddd-template/infra/database"
	"ddd-template/infra/xgin"
	"go.uber.org/zap"
)

import (
	_ "ddd-template/docs"
)

// Injectors from wire.go:

func Init(cfg conf.Configs, log *zap.Logger) *restful.Rest {
	engine := xgin.NewEngine(cfg)
	db := database.NewDatabase(cfg)
	demoInterface := repositry.NewDemoDependencyImpl(db, log)
	demoService := service.NewDemoService(demoInterface, log)
	demoServer := application.NewDemoServer(demoService, log)
	demoCtl := restful.NewDemoCtl(demoServer, log)
	rest := restful.NewRest(cfg, engine, demoCtl, log)
	return rest
}
