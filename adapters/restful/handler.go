package restful

import (
	"ddd-template/app/serve"
	"ddd-template/common/schema"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DemoHandler interface {
	SayHello(ctx *gin.Context)
	DoHandles(g gin.IRouter)
}

type DemoHandlerImpl struct {
	server serve.DemoServer
	log    *zap.Logger
}

//
// NewDemoCtl
// #Description: demo controller
// #param server app.DemoServer
// #param logger *zap.Logger
func NewDemoCtl(server serve.DemoServer, logger *zap.Logger) DemoHandler {
	handler := &DemoHandlerImpl{server: server, log: logger}
	return handler
}

func (r *DemoHandlerImpl) DoHandles(g gin.IRouter) {
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
// @Success 200 {object} Result{code=int,data=schema.DemoInfo}
func (r *DemoHandlerImpl) SayHello(ctx *gin.Context) {
	msg := ctx.Query("msg")
	var (
		err error
		res *schema.DemoInfo
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
