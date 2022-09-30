package delivery

import (
	"ddd-template/internal/schema"
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
func (h server) AddStudent(ctx *fiber.Ctx) (err error) {
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
	if err = h.app.Mutations.Student.AddStudent(ctx.UserContext(), form); err != nil {
		return
	}
	return resp.Success(ctx, nil)
}
