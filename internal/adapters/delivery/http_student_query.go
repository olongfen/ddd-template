package delivery

import (
	"ddd-template/internal/schema"
	"ddd-template/pkg/error_i18n"
	"ddd-template/pkg/response"
	"ddd-template/pkg/scontext"
	"github.com/gofiber/fiber/v2"
)

// GetStudent
// @Id student one
// @tags students
// @Summary get student one record
// @Description get
// @Param id  path int true "id"
// @router /api/v1/students/{id} [get]
// @Success 200 {object} response.Response{data=schema.StudentResp}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (h server) GetStudent(ctx *fiber.Ctx) (err error) {
	var (
		id       int
		language = scontext.GetLanguage(ctx.UserContext())
	)
	if id, err = ctx.ParamsInt("id"); err != nil {
		err = error_i18n.NewError(error_i18n.IllegalParameter, language)
		return
	}
	student, err := h.app.Queries.Student.GetStudent(ctx.UserContext(), id)
	if err != nil {
		return err
	}
	return response.NewResponse(language).Success(ctx, student)
}

// QueryStudents
// @Id query students
// @tags students
// @Summary query students
// @Description get
// @Param {} query schema.StudentsQuery true "query struct"
// @router /api/v1/students [get]
// @Success 200 {object} response.Response{code=int,data=schema.StudentsResp}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (h server) QueryStudents(ctx *fiber.Ctx) (err error) {
	var (
		query    = new(schema.StudentsQuery)
		language = scontext.GetLanguage(ctx.UserContext())
		resp     = response.NewResponse(language)
		data     schema.StudentsResp
		page     *schema.Pagination
	)
	if err = ctx.QueryParser(query); err != nil {
		return
	}
	if errors := schema.ValidateForm(query, language); len(errors) != 0 {
		err = setValidateDetail(resp, errors, language)
		ctx.SetUserContext(scontext.SetErrorsContext(ctx.UserContext(), errors))
		return
	}
	if data, page, err = h.app.Queries.Student.QueryStudents(ctx.UserContext(), *query); err != nil {
		return
	}

	return resp.Success(ctx, map[string]any{
		"list":       data,
		"pagination": page,
	})
}

func setValidateDetail(resp *response.Response, errors map[string]*schema.Error, lange string) error {
	resp.SetErrors(errors)
	return error_i18n.NewError(error_i18n.IllegalParameter, lange)
}
