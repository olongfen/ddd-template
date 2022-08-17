//go:build wireinject
// +build wireinject

package main

import (
	"ddd-template/internal/app"
	"ddd-template/internal/initialization"
	"ddd-template/internal/service/delivery/xfiber"
	"ddd-template/internal/service/repository"
	"ddd-template/internal/service/usecase"
	"github.com/google/wire"
)

func NewServer(confPath string) error {
	panic(wire.Build(app.Set, xfiber.Set, usecase.Set,
		repository.Set, initialization.Set))
}
