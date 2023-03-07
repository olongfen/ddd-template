package handler

import (
	app "ddd-template/internal/application"
	"github.com/gofiber/fiber/v2"
	"github.com/olongfen/toolkit/response"
	"github.com/olongfen/toolkit/scontext"
)

// demo demo handler
type demo struct {
	app *app.Application
}

// hello
// @tags Demo
// @Summary hello world
// @Description
// @router /api/v1/demo [get]
// @Success 200 {object} response.Response{}
// @Security ApiKeyAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (h *demo) hello(c *fiber.Ctx) error {
	var (
		ctx  = c.UserContext()
		lan  = scontext.GetLanguage(ctx)
		resp = response.NewResponse(lan)
	)
	return resp.Success(c, h.app.Query().Demo().Hello(ctx))

}
