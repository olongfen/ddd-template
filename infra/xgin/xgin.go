package xgin

import (
	"ddd-template/infra/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewEngine)

//
// NewEngine
// #Description:
// #param cfg conf.Server
// #return *Engine
func NewEngine(cfg conf.Configs) *gin.Engine {
	e := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	e.Use(cors.New(corsCfg))
	return e
}
