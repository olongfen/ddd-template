// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-template/adapters/repositry"
	"ddd-template/adapters/restful"
	"ddd-template/adapters/rpcx"
	"ddd-template/app"
	"ddd-template/app/serve"
	"ddd-template/common/conf"
	"go.uber.org/zap"
)

import (
	_ "ddd-template/docs"
)

// Injectors from wire.go:

func Init(cfg conf.Configs, log *zap.Logger) *app.Application {
	httpServer := restful.NewHttpServer(cfg)
	db := repositry.NewDatabase(cfg)
	demoInterface := repositry.NewDemoDependencyImpl(db, log)
	demoServer := serve.NewDemoServer(demoInterface, log)
	demoHandler := restful.NewDemoCtl(demoServer, log)
	greeterServer := rpcx.NewDemoGrpcServer(demoServer, log)
	grpcServer := rpcx.NewGrpc(greeterServer)
	application := app.NewApp(httpServer, demoHandler, grpcServer, log)
	return application
}
