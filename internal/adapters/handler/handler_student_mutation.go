package handler

import (
	"ddd-template/internal/application/schema"
	"ddd-template/pkg/error_i18n"
	"ddd-template/pkg/response"
	"ddd-template/pkg/scontext"
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
func (s handler) AddStudent(ctx *fiber.Ctx) (err error) {
	var (
		form     = new(schema.StudentAddForm)
		language = scontext.GetLanguage(ctx.UserContext())
		resp     = response.NewResponse(language)
	)
	if err = ctx.BodyParser(form); err != nil {
		return
	}
	// 验证表单
	if errs := schema.ValidateForm(form, language); len(errs) != 0 {
		ctx.SetUserContext(scontext.SetErrorsContext(ctx.UserContext(), errs))
		err = error_i18n.NewError(error_i18n.IllegalParameter, language)
		return
	}
	if err = s.app.Mutations.Student.AddStudent(ctx.UserContext(), form); err != nil {
		return
	}
	return resp.Success(ctx, nil)
}

// UpStudent
// @Id  student update one
// @tags students
// @Summary 更新
// @Description 通过id更新一条学生记录
// @Param id path int true "id"
// @Param {} body schema.StudentUpForm true "update form"
// @router /api/v1/students/{id} [put]
// @Success 200 {object} response.Response{}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (s handler) UpStudent(ctx *fiber.Ctx) (err error) {
	var (
		form     = new(schema.StudentUpForm)
		language = scontext.GetLanguage(ctx.UserContext())
		resp     = response.NewResponse(language)
		id       int
	)

	if id, err = ctx.ParamsInt("id"); err != nil || id <= 0 {
		err = error_i18n.NewError(error_i18n.IllegalParameter, language)
		// 定义逻辑错误
		ctx.SetUserContext(scontext.SetErrorsContext(ctx.UserContext(), map[string]*schema.Error{
			"id": {
				Failed: "id",
				Tag:    "id",
				Value:  id,
				Detail: err.Error(),
			},
		}))
		return
	}

	if err = ctx.BodyParser(form); err != nil {
		return
	}

	if errs := schema.ValidateForm(form, language); len(errs) != 0 {
		ctx.SetUserContext(scontext.SetErrorsContext(ctx.UserContext(), errs))
		err = error_i18n.NewError(error_i18n.IllegalParameter, language)
		return
	}

	if err = s.app.Mutations.Student.UpStudent(ctx.UserContext(), uint(id), form); err != nil {
		return
	}

	return resp.Success(ctx, nil)
}
