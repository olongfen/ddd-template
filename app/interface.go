package app

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type DemoHandler interface {
	SayHello(ctx *gin.Context)
	Handles(g gin.IRouter)
}

type HttpServer interface {
	gin.IRouter
	Run(addr ...string) error
}

type GrpcServer interface {
	Run(addr string) error
	SetOptions(opts ...grpc.ServerOption) GrpcServer
}
