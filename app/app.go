package app

import (
	"ddd-template/common/conf"
	_ "ddd-template/docs"
	"fmt"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"sync"
)

var ProviderSet = wire.NewSet(NewApp)

type Application struct {
	http        HttpServer
	grpc        GrpcServer
	demoHandler DemoHandler
	log         *zap.Logger
}

func NewApp(rest HttpServer, demoHandler DemoHandler, rpc GrpcServer, logger *zap.Logger) *Application {
	return &Application{http: rest, demoHandler: demoHandler, grpc: rpc, log: logger}
}

//
// Handles
// #Description: 所有的接口统一在app这里处理
// #receiver a *Application
// #param basePath string
// #return *Application
func (a *Application) Handles(basePath string) *Application {
	group := a.http.Group(basePath)
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	a.demoHandler.Handles(group)
	return a
}

//
// Run
// #Description: run server
// #receiver a *Application
// #param cfg conf.Server
func (a *Application) Run(cfg conf.Server) {
	var (
		wg  = sync.WaitGroup{}
		err error
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		a.log.Sugar().Infof("http server run in: %s", cfg.Http.Addr)
		if err = a.http.Run(cfg.Http.Addr); err != nil {
			a.log.Fatal(err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		addr := fmt.Sprintf("%s:%d", cfg.GRpc.Host, cfg.GRpc.Port)
		a.log.Sugar().Infof("grpc server run in: %s", addr)
		if err = a.grpc.Run(addr); err != nil {
			a.log.Fatal(err.Error())
		}
	}()
	wg.Wait()
	return
}
