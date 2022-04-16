package restful

import (
	v1 "ddd-template/api/v1"
	_ "ddd-template/docs"
	"ddd-template/internal/app"
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/xlog"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type HTTPServer struct {
	engine *gin.Engine
	cfg    conf.HTTP
	demo   v1.GreeterServer
}

func NewHTTPServer(demo v1.GreeterServer, cfg *conf.Configs) app.HttpServer {
	h := new(HTTPServer)
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	h.cfg = cfg.Server.Http
	h.demo = demo
	h.engine = gin.Default()
	return h
}

func (h *HTTPServer) Run() error {
	h.engine.Use(corsHandler())
	group := h.engine.Group("/api/v1")
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	xlog.Log.Sugar().Infof("http server run in: %s", h.cfg.Addr)
	group1 := group.Group("/demo")
	v1.RegisterGreeterHTTPServer(group1, h.demo)
	return h.engine.Run(h.cfg.Addr)
}
