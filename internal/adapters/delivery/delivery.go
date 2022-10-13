package delivery

import (
	app "ddd-template/internal/application"
	"ddd-template/internal/ports/controller"
)

type server struct {
	app app.Application
}

func (s server) Cleanup() func() {
	return s.app.Cleanup
}

func NewServer(app app.Application) controller.Server {
	return &server{app: app}
}
