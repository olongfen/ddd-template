package app

import "context"

type HttpServer interface {
	Run() error
}

type GrpcServer interface {
	Run() error
}

type ITransaction interface {
	ExecTx(ctx context.Context, fc func(ctx context.Context) error) error
}
