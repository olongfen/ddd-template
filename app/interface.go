package app

import (
	"github.com/gin-gonic/gin"
)

type DemoHandler interface {
	SayHello(ctx *gin.Context)
	Handles(g gin.IRouter)
}

type HttpServer interface {
	gin.IRouter
	Run(addr ...string) error
}
