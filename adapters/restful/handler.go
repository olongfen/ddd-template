package restful

import (
	"ddd-template/app"
	"ddd-template/app/dto"
	"ddd-template/app/serve"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DemoHandler struct {
	server serve.DemoServer
	log    *zap.Logger
}

//
// NewDemoCtl
// #Description: demo controller
// #param server app.DemoServer
// #param logger *zap.Logger
func NewDemoCtl(server serve.DemoServer, logger *zap.Logger) app.DemoHandler {
	handler := &DemoHandler{server: server, log: logger}
	return handler
}

func (r *DemoHandler) Handles(g gin.IRouter) {
	group := g.Group("/demo")
	group.GET("/", r.SayHello)
}

// SayHello
// @Tags Demo
// @Summary say hello
// @Description sends a string msg
// @Accept application/json
// @Produce application/json
// @Param msg query string false "message"
// @Router /api/v1/demo  [get]
// @Success 200 {object} Result{code=int,data=dto.DemoInfo}
func (r *DemoHandler) SayHello(ctx *gin.Context) {
	msg := ctx.Query("msg")
	var (
		err error
		res *dto.DemoInfo
	)
	defer func() {
		if err != nil {
			SendFail(ctx, err)
		} else {
			SendSuccess(ctx, res)
		}
	}()
	if res, err = r.server.SayHello(ctx.Request.Context(), msg); err != nil {
		return
	}
}
