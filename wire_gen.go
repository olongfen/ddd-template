// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-template/internal/adapters/repository"
	"ddd-template/internal/application"
	"ddd-template/internal/application/mutation"
	"ddd-template/internal/application/query"
	"ddd-template/internal/ports/controller"
	"ddd-template/internal/ports/controller/handler"
	"ddd-template/internal/ports/graph"
	"ddd-template/internal/rely"
	"ddd-template/internal/service"
	"github.com/olongfen/toolkit/db_data"
)

// Injectors from wire.go:

func NewServer(configFile2 string) (*service.Server, func()) {
	configs := rely.InitConfigs(configFile2)
	logger := rely.NewLogger(configs)
	db := rely.InitDBConnect(configs, logger)
	dbData, cleanup := db_data.NewData(db, logger)
	iDemoRepo := repository.NewDemo(dbData)
	iDemoService := mutation.NewDemo(iDemoRepo)
	mutationMutation := mutation.SetMutation(iDemoService)
	queryIDemoService := query.NewDemo(iDemoRepo)
	queryQuery := query.SetQuery(queryIDemoService)
	application := app.NewApplication(mutationMutation, queryQuery)
	handlerHandler := handler.NewHandler(application)
	resolver := graph.NewResolver(application, logger)
	middleware := controller.NewMiddleware(logger)
	httpServer, cleanup2 := controller.NewHTTPServer(handlerHandler, resolver, middleware, configs, logger)
	server := service.NewServer(httpServer)
	return server, func() {
		cleanup2()
		cleanup()
	}
}
