package xfiber

import (
	"ddd-template/internal/common/response"
	"ddd-template/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type DemoHandler struct {
	us domain.IDemoUsecase
}

func NewDemoHandler(us domain.IDemoUsecase) *DemoHandler {
	return &DemoHandler{us: us}
}

func (d *DemoHandler) Handler(f fiber.Router) {
	demo := f.Group("/demo")
	demo.Get("/hello", d.SayHello)
}

// SayHello
// @tags demo
// @Summary hello
// @Description hello
// @router /api/v1/demo/hello [get]
// @Success 200 {object} response.Response
// @Failure 404 {object} string
// @Failure 500 {object} string
func (d *DemoHandler) SayHello(ctx *fiber.Ctx) (err error) {
	return response.FiberRespSuccessFunc(ctx, "hello")
}
