package usecase

import (
	"context"
	"ddd-template/internal/domain"
	"go.uber.org/zap"
)

type demoServerImpl struct {
	repo domain.IDemoRepo
	log  *zap.Logger
	tx   domain.ITransaction
}

//
// NewDemoServer
// #Description: new
// #param demo dependency.IDemoRepo
// #return IDemoUsecase
func NewDemoServer(demo domain.IDemoRepo, tx domain.ITransaction, logger *zap.Logger) domain.IDemoUsecase {
	return &demoServerImpl{demo, logger, tx}
}

//
// SayHello
// #Description: demo server use case
// #receiver s *demoServerImpl
// #param ctx context.Context
// #param msg string
// #return string
func (s *demoServerImpl) SayHello(ctx context.Context, msg string) (res *domain.Demo, err error) {
	var (
		data *domain.Demo
	)
	data = s.repo.SayHello(ctx, msg)
	res = data
	return

}
