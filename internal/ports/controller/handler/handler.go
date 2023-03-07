package handler

import (
	app "ddd-template/internal/application"
	"github.com/gofiber/fiber/v2"
)

// Handler http handler
type Handler struct {
	app *app.Application
}

// NewHandler new
func NewHandler(app *app.Application) *Handler {
	return &Handler{app: app}
}

// Process  handler process
func (h *Handler) Process(group fiber.Router) *Handler {
	d := &demo{app: h.app}
	dRouter := group.Group("/demo")
	dRouter.Get("/", d.hello)
	return h
}
