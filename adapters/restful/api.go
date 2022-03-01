package restful

import (
	"ddd-template/application"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DemoCtl struct {
	server application.DemoServer
	log    *zap.Logger
}

//
// NewDemoCtl
// #Description: demo controller
// #param server application.DemoServer
// #param logger *zap.Logger
// #return *DemoCtl
func NewDemoCtl(server application.DemoServer, logger *zap.Logger) *DemoCtl {
	return &DemoCtl{server: server, log: logger}
}

// SayHello
// @Tags Demo
// @Summary say hello
// @Description sends a string msg
// @Accept application/json
// @Produce application/json
// @Param msg query string false "message"
// @Router /api/v1/demo  [get]
// @Success 200 {object} Result{code=int,data=string}
func (r *DemoCtl) SayHello(ctx *gin.Context) {
	msg := ctx.Query("msg")
	SendSuccess(ctx, r.server.SayHello(ctx.Request.Context(), msg))
}
