// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-template/internal/app"
	"ddd-template/internal/initialization"
	"ddd-template/internal/service/delivery/xfiber"
	"ddd-template/internal/service/repository"
	"ddd-template/internal/service/usecase"
)

// Injectors from wire.go:

func NewServer(confPath string) error {
	configs := initialization.InitConf(confPath)
	logger := initialization.InitLog(configs)
	db := repository.NewDB(configs, logger)
	data := repository.NewData(db, logger)
	iDemoRepo := repository.NewDemoDependency(data, logger)
	iTransaction := repository.NewTransaction(data)
	iDemoUsecase := usecase.NewDemoServer(iDemoRepo, iTransaction, logger)
	demoHandler := xfiber.NewDemoHandler(iDemoUsecase)
	error2 := app.NewApp(configs, logger, demoHandler)
	return error2
}
