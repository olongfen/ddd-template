package xfiber

import (
	"ddd-template/internal/common/response"
	"ddd-template/internal/common/utils"
	xi18n2 "ddd-template/internal/common/xi18n"
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type DemoHandler struct {
	us   domain.IDemoUseCase
	i18n *i18n.Bundle
}

func NewDemoHandler(us domain.IDemoUseCase, i18n *i18n.Bundle) *DemoHandler {
	return &DemoHandler{us: us, i18n: i18n}
}

func (d *DemoHandler) Handler(f fiber.Router) {
	demo := f.Group("/demo")
	demo.Get("/:id", d.Get)
}

// Get
// @Id demo-get-one
// @tags demo
// @Summary hello
// @Description hello
// @Param id path int true "id"
// @router /demo/{id} [get]
// @Success 200 {object} response.Response{data=schema.DemoResp}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (d *DemoHandler) Get(ctx *fiber.Ctx) (err error) {
	var (
		hello *domain.Demo
		data  schema.DemoResp
		id    int
	)
	lo := i18n.NewLocalizer(d.i18n, utils.GetLanguage(ctx.UserContext()))
	if id, err = ctx.ParamsInt("id"); err != nil {
		err = xi18n2.NewError(lo.MustLocalize(&i18n.LocalizeConfig{MessageID: xi18n2.IllegalParameter})).WithError(err)
		return
	}
	hello, err = d.us.Get(ctx.UserContext(), uint(id))
	if err != nil {
		return err
	}
	if err = utils.Copier(hello, &data); err != nil {
		return
	}
	return response.SuccessFunc(ctx, data)
}
