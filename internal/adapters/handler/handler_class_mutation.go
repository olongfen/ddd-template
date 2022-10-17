package handler

import (
	"ddd-template/internal/schema"
	"ddd-template/pkg/error_i18n"
	"ddd-template/pkg/response"
	"ddd-template/pkg/scontext"
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
func (s handler) AddClass(ctx *fiber.Ctx) (err error) {
	var (
		form     = new(schema.ClassAddForm)
		language = scontext.GetLanguage(ctx.UserContext())
		resp     = response.NewResponse(language)
	)
	if err = ctx.BodyParser(form); err != nil {
		return
	}
	// 验证表单
	if errs := schema.ValidateForm(form, language); len(errs) != 0 {
		err = error_i18n.NewError(error_i18n.IllegalParameter, language)
		ctx.SetUserContext(scontext.SetErrorsContext(ctx.UserContext(), errs))
		return
	}
	if err = s.app.Mutations.Class.AddClass(ctx.UserContext(), form); err != nil {
		return
	}
	return resp.Success(ctx, nil)
}

// UpClass
//
// @Id update class
// @tags classes
// @Summary update class
// @Description update class by id
// @Param {} body schema.ClassUpForm true "form"
// @Param id path int  true "id"
// @router /api/v1/classes/{id} [put]
// @Success 200 {object} response.Response{}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (s handler) UpClass(ctx *fiber.Ctx) (err error) {
	var (
		form = new(schema.ClassUpForm)
		lan  = scontext.GetLanguage(ctx.UserContext())
		resp = response.NewResponse(lan)
	)
	if err = ctx.ParamsParser(form); err != nil {
		return
	}
	if err = ctx.BodyParser(form); err != nil {
		return
	}
	// 验证表单
	if errs := schema.ValidateForm(form, lan); len(errs) != 0 {
		err = error_i18n.NewError(error_i18n.IllegalParameter, lan)
		ctx.SetUserContext(scontext.SetErrorsContext(ctx.UserContext(), errs))
		return
	}

	if err = s.app.Mutations.Class.UpClass(ctx.UserContext(), form); err != nil {
		return
	}

	return resp.Success(ctx, nil)
}
