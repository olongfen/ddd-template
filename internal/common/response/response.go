package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

type Func func(ctx *fiber.Ctx, data interface{}, status ...int) error

func SetResponseSuccessFunc(fc Func) {
	RespSuccessFunc = fc
}
func SetResponseFailFunc(fc Func) {
	RespFailFunc = fc
}

var RespSuccessFunc Func = func(ctx *fiber.Ctx, data interface{}, status ...int) error {
	var (
		code = 200
	)
	if len(status) > 0 {
		code = status[0]
	}
	return ctx.Status(code).JSON(&Response{Code: 0, Data: data})
}

var RespFailFunc Func = func(ctx *fiber.Ctx, msg interface{}, status ...int) error {
	var (
		code = 200
	)
	if len(status) > 0 {
		code = status[0]
	}
	return ctx.Status(code).JSON(&Response{Code: -1, Message: msg})
}
