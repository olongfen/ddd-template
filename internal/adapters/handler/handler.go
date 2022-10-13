package handler

import (
	app "ddd-template/internal/application"
	"ddd-template/internal/ports/controller"
)

type handler struct {
	app app.Application
}

func (s handler) Cleanup() func() {
	return s.app.Cleanup
}

func NewServer(app app.Application) controller.Server {
	return &handler{app: app}
}
