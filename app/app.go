package app

import (
	_ "ddd-template/docs"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var ProviderSet = wire.NewSet(NewApp)

type Application struct {
	http        HttpServer
	demoHandler DemoHandler
}

func NewApp(rest HttpServer, demoHandler DemoHandler) *Application {
	return &Application{http: rest, demoHandler: demoHandler}
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

func (a *Application) Run(addr ...string) error {
	return a.http.Run(addr...)
}
