package application

import (
	"context"
	"ddd-template/domain/service"
	"go.uber.org/zap"
)

type DemoServer interface {
	SayHello(ctx context.Context, msg string) string
}

type DemoServerImpl struct {
	demoService *service.DemoService
	log         *zap.Logger
}

//
// NewDemoServer
// #Description: new
// #param demo *service.DemoService
// #return DemoServer
func NewDemoServer(demo *service.DemoService, logger *zap.Logger) DemoServer {
	return &DemoServerImpl{demo, logger}
}

//
// SayHello
// #Description: demo server use case
// #receiver s *DemoServerImpl
// #param ctx context.Context
// #param msg string
// #return string
func (s *DemoServerImpl) SayHello(ctx context.Context, msg string) string {
	return s.demoService.SayHello(ctx, msg)

}
