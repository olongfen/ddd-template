package handler

import (
	"ddd-template/internal/application/schema"
	"ddd-template/pkg/error_i18n"
	"ddd-template/pkg/response"
	"ddd-template/pkg/scontext"
	"github.com/gofiber/fiber/v2"
)

// QueryClasses
// @Id query class
// @tags classes
// @Summary 查询class
// @Description 分页查询class
// @Param   order  query     []string   false  "string order collection"  collectionFormat(multi)
// @Param   sort  query     []string   false  "string sort collection"  collectionFormat(multi)
// @Param {} query schema.ClassQueryReq true "query"
// @router /api/v1/classes [get]
// @Success 200 {object} response.Response{data=schema.ClassQueryResp}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (s handler) QueryClasses(ctx *fiber.Ctx) (err error) {
	var (
		query = new(schema.ClassQueryReq)
		lan   = scontext.GetLanguage(ctx.UserContext())
		resp  = response.NewResponse(lan)
		data  schema.ClassRespList
		pag   *schema.Pagination
	)
	if err = ctx.QueryParser(query); err != nil {
		return
	}
	if errs := schema.ValidateForm(query, lan); len(errs) != 0 {
		err = error_i18n.NewError(error_i18n.IllegalParameter, lan)
		ctx.SetUserContext(scontext.SetErrorsContext(ctx.UserContext(), errs))
		return
	}
	if data, pag, err = s.app.Queries.Class.QueryClasses(ctx.UserContext(), query); err != nil {
		return
	}
	return resp.Success(ctx, schema.ClassQueryResp{
		List:       data,
		Pagination: pag,
	})
}

// GetClass
// @Id get class one
// @tags classes
// @Summary 获取一条class
// @Description 通过id获取
// @Param id path int true "id"
// @router /api/v1/classes/{id} [get]
// @Success 200 {object} response.Response{data=schema.ClassResp}
// @Security BearerAuth
// @Failure 404 {object} string
// @Failure 500 {object} string
func (s handler) GetClass(ctx *fiber.Ctx) (err error) {
	var (
		lan  = scontext.GetLanguage(ctx.UserContext())
		resp = response.NewResponse(lan)
		data *schema.ClassResp
		id   int
	)
	if id, err = ctx.ParamsInt("id"); err != nil {
		err = error_i18n.NewError(error_i18n.IllegalParameter, lan)
		return
	}
	if data, err = s.app.Queries.Class.GetClass(ctx.UserContext(), id); err != nil {
		return
	}
	return resp.Success(ctx, data)
}
