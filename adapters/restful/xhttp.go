package restful

import (
	"ddd-template/application"
	"ddd-template/infra/conf"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewRest)

type Rest struct {
	engine *gin.Engine
	addr   string
	app    application.Server
	log    *zap.Logger
}

//
// NewRest
// #Description: new rest server
// #param cfg *conf.Configs
// #param e *gin.Engine
// #param app application.Server
// #return *Rest
func NewRest(cfg conf.Configs, e *gin.Engine, app application.Server, logger *zap.Logger) *Rest {
	return &Rest{engine: e, addr: cfg.Server.Http.Addr, app: app, log: logger}
}

//
// Router
// #Description: set restful router
// #receiver r *Rest
// #return *Rest
func (r *Rest) Router() *Rest {
	group := r.engine.Group("/api/v1")
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group.GET("/demo", r.app.SayHello)
	return r
}

//
// InjectMiddleware
// #Description: inject middleware
// #receiver r *Rest
// #param middleware ...gin.HandlerFunc
// #return *Rest
func (r *Rest) InjectMiddleware(middleware ...gin.HandlerFunc) *Rest {
	r.engine.Use(middleware...)
	return r
}

//
// Start
// #Description: start restful server
// #receiver r *Rest
// #return error
func (r *Rest) Start() error {
	return r.engine.Run(r.addr)
}
