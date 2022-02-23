package application

import (
	"ddd-template/application/errorx"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`    // 业务码 等于0时表示业务逻辑执行成功,其他表示业务逻辑失败
	Data    interface{} `json:"data"`    // 数据对象
	Message string      `json:"message"` // 信息
}

//
// SendSuccess
// #Description: response success
// #param ctx *gin.Context
// #param data interface{}
func SendSuccess(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(200, Result{Code: 0, Data: data, Message: "success"})
}

//
// SendFail
// #Description: response fail
// #param ctx *gin.Context
// #param err errorx.BusinessError
func SendFail(ctx *gin.Context, err errorx.BusinessError) {
	ctx.AbortWithStatusJSON(err.HTTPCode(), Result{Code: err.BusinessCode(), Message: err.Message()})
}
