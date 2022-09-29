package controller

import (
	"ddd-template/internal/schema"
	"ddd-template/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// AddClass
// @Id class add one
// @tags classes
// @Summary add class one record
// @Description add
// @Param {}  body schema.ClassAddForm true "form"
// @router /api/v1/classes [post]
// @Success 200 {object} response.Response{}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (h HttpServer) AddClass(ctx *fiber.Ctx) (err error) {
	var (
		form = new(schema.ClassAddForm)
	)
	if err = ctx.BodyParser(form); err != nil {
		return
	}
	if err = h.app.Mutations.Class.AddClass(ctx.UserContext(), form); err != nil {
		return
	}
	return response.SuccessHandler(ctx, nil)
}
