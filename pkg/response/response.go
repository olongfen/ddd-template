package response

import (
	"ddd-template/pkg/error_i18n"
	"ddd-template/pkg/scontext"
	"ddd-template/pkg/xlog"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	status int
	//
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	Language   string      `json:"language"`
	Pagination interface{} `json:"pagination"`
	Errors     interface{} `json:"errors"`
}

func NewResponse(language string) *Response {
	return &Response{status: http.StatusOK, Language: language}
}

func (r *Response) SetPagination(pagination interface{}) *Response {
	r.Pagination = pagination
	return r
}

func (r *Response) SetErrors(errs interface{}) *Response {
	r.Errors = errs
	return r
}

func (r *Response) SetMessage(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) Success(ctx *fiber.Ctx, data interface{}) error {
	r.Data = data
	if r.Message == "" {
		r.Message = "success"
	}
	return ctx.Status(r.status).JSON(r)
}

var ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		status = e.Code
	}
	if e, ok := err.(error_i18n.BizError); ok {
		xlog.Log.Error("Business Error", zap.Error(e.StackError()))
		status = fiber.StatusOK
	}
	userCtx := ctx.UserContext()
	resp := NewResponse(scontext.GetLanguage(userCtx))
	resp.SetErrors(scontext.GetErrorsContext(userCtx))
	resp.Message = err.Error()
	return ctx.Status(status).JSON(resp)
}
