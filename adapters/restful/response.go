package restful

import (
	"ddd-template/app/errorx"
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
// #param err error
func SendFail(ctx *gin.Context, err error) {
	switch err.(type) {
	case errorx.BusinessError:
		e := err.(errorx.BusinessError)
		ctx.AbortWithStatusJSON(e.HTTPCode(), Result{Code: e.BusinessCode(), Message: e.Error()})
	default:
		ctx.AbortWithStatusJSON(200, Result{Code: -1, Message: err.Error()})
	}

}
