package handler

import (
	app "ddd-template/internal/application"
)

type Handler struct {
	app *app.Application
}

func NewHandler(app *app.Application) *Handler {
	return &Handler{app: app}
}
