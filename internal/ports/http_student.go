package ports

import (
	"ddd-template/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// GetStudent
// @Id student one
// @tags student
// @Summary get student one record
// @Description get
// @Param uuid  path string true "uuid4"
// @router /api/v1/student/{uuid} [get]
// @Success 200 {object} response.Response{}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (h HttpServer) GetStudent(ctx *fiber.Ctx) (err error) {
	return response.SuccessHandler(ctx, ctx.Params("uuid"))
}
