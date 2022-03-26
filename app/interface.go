package app

type HttpServer interface {
	Run(basePath, addr string) error
}

type GrpcServer interface {
	Run(addr string) error
}
