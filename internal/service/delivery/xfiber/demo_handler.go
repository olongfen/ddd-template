package xfiber

import (
	"ddd-template/internal/common/response"
	"ddd-template/internal/common/utils"
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
	"github.com/gofiber/fiber/v2"
)

type DemoHandler struct {
	us domain.IDemoUseCase
}

func NewDemoHandler(us domain.IDemoUseCase) *DemoHandler {
	return &DemoHandler{us: us}
}

func (d *DemoHandler) Handler(f fiber.Router) {
	demo := f.Group("/demo")
	demo.Get("/:id", d.Get)
}

// Get
// @tags demo
// @Summary hello
// @Description hello
// @Param id path int true "id"
// @router /api/v1/demo/{id} [get]
// @Success 200 {object} response.Response{data=schema.DemoResp}
// @Failure 404 {object} string
// @Failure 500 {object} string
func (d *DemoHandler) Get(ctx *fiber.Ctx) (err error) {
	var (
		hello *domain.Demo
		data  schema.DemoResp
		id    int
	)
	defer func() {
		if err != nil {
			err = response.RespFailFunc(ctx, err.Error())
		} else {
			err = response.RespSuccessFunc(ctx, data)
		}
	}()
	if id, err = ctx.ParamsInt("id"); err != nil {
		return
	}
	hello, err = d.us.Get(ctx.UserContext(), uint(id))
	if err != nil {
		return err
	}
	if err = utils.Copier(hello, &data); err != nil {
		return
	}
	return
}