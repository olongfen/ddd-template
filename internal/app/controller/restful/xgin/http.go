package xgin

import (
	_ "ddd-template/api"
	v1 "ddd-template/api/v1"
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

func (h *HTTPServer) Handlers() app.HTTPServer {
	h.engine.Use(corsHandler(), logger())
	h.engine.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	xlog.Log.Sugar().Infof("http server run in: %s", h.cfg.Addr)
	group1 := h.engine.Group("/")
	v1.RegisterGreeterGinHTTPServer(group1, h.demo)
	return h
}

func NewHTTPServer(demo v1.GreeterServer, cfg *conf.Configs) app.HTTPServer {
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
	return h.engine.Run(h.cfg.Addr)
}
