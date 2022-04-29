//go:build wireinject
// +build wireinject

package main

import (
	"ddd-template/internal/adapters/repository"
	"ddd-template/internal/adapters/restful"
	"ddd-template/internal/adapters/rpcx"
	"ddd-template/internal/app"
	"ddd-template/internal/app/service"
	"ddd-template/internal/app/usecase"
	"ddd-template/internal/initialization"
	"github.com/google/wire"
)

func NewServer(confPath string) (*app.Application, error) {
	panic(wire.Build(app.Set, rpcx.Set, restful.Set, service.Set, usecase.Set,
		repository.Set, initialization.Set))
}
