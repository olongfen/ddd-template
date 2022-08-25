package response

import (
	"ddd-template/internal/common/errorx"
	"ddd-template/internal/common/xlog"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

var SuccessFunc = func(ctx *fiber.Ctx, data interface{}, status ...int) error {
	var (
		code = 200
	)
	if len(status) > 0 {
		code = status[0]
	}
	return ctx.Status(code).JSON(&Response{Code: 0, Data: data})
}

var ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	if e, ok := err.(errorx.BizError); ok {
		xlog.Log.Error("Business Error", zap.Error(e.StackError()))
		code = fiber.StatusOK
	}
	return ctx.Status(code).JSON(&Response{Code: -1, Message: err.Error()})
}
