// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-template/internal/app"
	"ddd-template/internal/app/controller/restful/xfiber"
	"ddd-template/internal/app/controller/rpcx"
	"ddd-template/internal/app/repository"
	"ddd-template/internal/app/service"
	"ddd-template/internal/app/usecase"
	"ddd-template/internal/initialization"
)

// Injectors from wire.go:

func NewServer(confPath string) (*app.Application, error) {
	configs := initialization.InitConf(confPath)
	logger := initialization.InitLog(configs)
	db := repository.NewDB(configs, logger)
	data := repository.NewData(db, logger)
	iDemoRepo := repository.NewDemoDependency(data, logger)
	iTransaction := repository.NewTransaction(data)
	iDemoUsecase := usecase.NewDemoServer(iDemoRepo, iTransaction, logger)
	greeterServer := service.NewDemoService(iDemoUsecase, logger)
	httpServer := xfiber.NewHTTPServer(greeterServer, configs)
	rpcServer := rpcx.NewGrpc(greeterServer, configs)
	application := app.NewApp(httpServer, rpcServer, logger)
	return application, nil
}
