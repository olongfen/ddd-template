package restful

import (
	"ddd-template/app"
	"ddd-template/common/conf"
	_ "ddd-template/docs"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var ProviderSet = wire.NewSet(NewDemoCtl, NewHTTPServerImpl)

type HTTPServerImpl struct {
	*gin.Engine
	demoHandler *DemoHandler
}

func NewHTTPServerImpl(cfg conf.Configs, demoCtl *DemoHandler) app.HttpServer {
	h := new(HTTPServerImpl)
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	h.demoHandler = demoCtl
	h.Engine = gin.Default()
	return h
}

func (h *HTTPServerImpl) Run(basePath, addr string) error {
	h.Engine.Use(corsHandler())
	group := h.Group(basePath)
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h.demoHandler.DoHandles(group)
	return h.Engine.Run(addr)
}
