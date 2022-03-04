package restful

import (
	"ddd-template/app"
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
func NewHttpServer() app.HttpServer {
	e := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	e.Use(cors.New(corsCfg))
	return e
}
