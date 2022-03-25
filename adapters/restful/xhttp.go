package restful

import (
	"ddd-template/app"
	"ddd-template/common/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewDemoCtl, NewHttpServer)

//
// NewHttpServer
// #Description: new rest server
// #param cfg *conf.Configs
// #param demoCtl app.DemoServer
// #return *Rest
func NewHttpServer(cfg conf.Configs) app.HttpServer {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	e := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	e.Use(cors.New(corsCfg))
	return e
}
