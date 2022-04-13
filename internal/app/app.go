package app

import (
	"go.uber.org/zap"
	"sync"
)

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
func (a *Application) Run() {
	var (
		wg  = sync.WaitGroup{}
		err error
	)
	wg.Add(2)
	go func() {
		defer wg.Done()

		if err = a.http.Run(); err != nil {
			a.log.Fatal(err.Error())
		}
	}()
	go func() {
		defer wg.Done()

		if err = a.grpc.Run(); err != nil {
			a.log.Fatal(err.Error())
		}
	}()
	wg.Wait()
	return
}
