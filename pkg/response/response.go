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
	Code     int         `json:"code"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Language string      `json:"language"`
	Errors   interface{} `json:"errors"`
}

func NewResponse(language string) *Response {
	return &Response{status: http.StatusOK, Language: language}
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
	status := fiber.StatusOK
	userCtx := ctx.UserContext()
	resp := NewResponse(scontext.GetLanguage(userCtx))
	switch err.(type) {
	case *fiber.Error:
		// 处理内部错误返回
		e := err.(*fiber.Error)
		xlog.Log.Error("HTTP Error", zap.Error(err))
		resp.Code = e.Code
		resp.Message = "failed"
	case error_i18n.BizError:
		// 处理自定义业务错误返回
		e := err.(error_i18n.BizError)
		xlog.Log.Error("Business Error", zap.Error(e.StackError()))
		resp.Code = e.Code()
		resp.Message = e.Error()
	case error_i18n.ValidateError:
		resp.SetErrors(err.(error_i18n.ValidateError))
		resp.Message = error_i18n.NewError(error_i18n.IllegalParameter, scontext.GetLanguage(userCtx)).Error()
	case error_i18n.DBErrorResponse:
		// 处理数据库错误返回
		var (
			m = map[string]string{}
		)
		e := err.(error_i18n.DBErrorResponse)
		for k, v := range e {
			resp.Code = v.Code()
			m[k] = v.Error()
		}
		resp.Message = "failed"
		resp.SetErrors(m)
	}
	/*	// 处理内部错误返回
		if e, ok := err.(*fiber.Error); ok {
			xlog.Log.Error("HTTP Error", zap.Error(err))
			resp.Code = e.Code
			resp.Message = "failed"
		}
		// 处理自定义业务错误返回
		if e, ok := err.(error_i18n.BizError); ok {
			xlog.Log.Error("Business Error", zap.Error(e.StackError()))
			resp.Code = e.Code()
			resp.Message = e.Error()
		}
		// 处理数据库错误返回
		if e, ok := err.(error_i18n.DBErrorResponse); ok {
			var (
				m = map[string]string{}
			)
			for k, v := range e {
				resp.Code = v.Code()
				m[k] = v.Error()
			}
			resp.Message = "failed"
			resp.SetErrors(m)
		}*/

	if resp.Errors == nil {
		resp.Errors = map[string]any{"error": err.Error()}
	}

	return ctx.Status(status).JSON(resp)
}
