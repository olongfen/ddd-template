package app

import "context"

type HTTPServer interface {
	Run() error
	Handlers() HTTPServer
}

type RPCServer interface {
	Run() error
	Stop()
	Handlers() RPCServer
}

type ITransaction interface {
	ExecTx(ctx context.Context, fc func(ctx context.Context) error) error
}
