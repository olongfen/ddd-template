package app

type HttpServer interface {
	Run() error
}

type GrpcServer interface {
	Run() error
}
