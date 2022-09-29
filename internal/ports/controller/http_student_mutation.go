package controller

import (
	"ddd-template/internal/schema"
	"ddd-template/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// AddStudent
// @Id student add one
// @tags students
// @Summary add student one record
// @Description add
// @Param {}  body schema.StudentAddForm true "form"
// @router /api/v1/students [post]
// @Success 200 {object} response.Response{}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (h HttpServer) AddStudent(ctx *fiber.Ctx) (err error) {
	var (
		form = new(schema.StudentAddForm)
	)
	if err = ctx.BodyParser(form); err != nil {
		return
	}
	if err = h.app.Mutations.Student.AddStudent(ctx.UserContext(), form); err != nil {
		return
	}
	return response.SuccessHandler(ctx, nil)
}
