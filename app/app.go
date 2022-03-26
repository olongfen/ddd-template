package app

import (
	"ddd-template/common/conf"
	_ "ddd-template/docs"
	"fmt"
	"github.com/google/wire"
	"go.uber.org/zap"
	"sync"
)

var ProviderSet = wire.NewSet(NewApp)

type Application struct {
	http HttpServer
	grpc GrpcServer
	log  *zap.Logger
}

func NewApp(rest HttpServer, rpc GrpcServer, logger *zap.Logger) *Application {
	return &Application{http: rest, grpc: rpc, log: logger}
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
		if err = a.http.Run("/api/v1", cfg.Http.Addr); err != nil {
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
