package serve

import (
	"context"
	"ddd-template/app/dto"
	"ddd-template/domain/dependency"
	"ddd-template/domain/entities"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewDemoServer)

type DemoServer interface {
	SayHello(ctx context.Context, msg string) (*dto.DemoInfo, error)
}
type demoServerImpl struct {
	repo dependency.DemoInterface
	log  *zap.Logger
}

//
// NewDemoServer
// #Description: new
// #param demo dependency.DemoInterface
// #return DemoServer
func NewDemoServer(demo dependency.DemoInterface, logger *zap.Logger) DemoServer {
	return &demoServerImpl{demo, logger}
}

//
// SayHello
// #Description: demo server use case
// #receiver s *demoServerImpl
// #param ctx context.Context
// #param msg string
// #return string
func (s *demoServerImpl) SayHello(ctx context.Context, msg string) (res *dto.DemoInfo, err error) {
	var (
		data *entities.Demo
	)
	data = s.repo.SayHello(ctx, msg)
	if res, err = dto.DemoEnt2Dto(*data); err != nil {
		s.log.Error(err.Error())
		return
	}
	return

}
