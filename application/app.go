package application

import (
	"ddd-template/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewServer)

type Server interface {
	SayHello(ctx *gin.Context)
}

type ServerImpl struct {
	demoService *service.DemoService
	log         *zap.Logger
}

//
// NewServer
// #Description: new
// #param demo *service.DemoService
// #return Server
func NewServer(demo *service.DemoService, logger *zap.Logger) Server {
	return &ServerImpl{demo, logger}
}
