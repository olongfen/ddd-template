package app

import (
	"go.uber.org/zap"
	"sync"
)

type Application struct {
	http HTTPServer
	grpc RPCServer
	log  *zap.Logger
}

func NewApp(rest HTTPServer, rpc RPCServer, logger *zap.Logger) *Application {
	return &Application{http: rest, grpc: rpc, log: logger}
}

func (a *Application) Run() {
	var (
		wg  = sync.WaitGroup{}
		err error
	)
	wg.Add(2)
	go func() {
		defer wg.Done()

		if err = a.http.Handlers().Run(); err != nil {
			a.log.Fatal(err.Error())
		}
	}()
	go func() {
		defer wg.Done()

		if err = a.grpc.Handlers().Run(); err != nil {
			a.log.Fatal(err.Error())
		}
	}()
	wg.Wait()
	return
}

func (a *Application) Close() error {
	a.grpc.Stop()
	return nil
}
