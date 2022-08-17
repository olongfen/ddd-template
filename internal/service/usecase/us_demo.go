package usecase

import (
	"context"
	"ddd-template/internal/app"
	"ddd-template/internal/domain"
	"go.uber.org/zap"
)

type demoServerImpl struct {
	repo domain.IDemoRepo
	log  *zap.Logger
	tx   app.ITransaction
}

func NewDemoServer(demo domain.IDemoRepo, tx app.ITransaction, logger *zap.Logger) domain.IDemoUsecase {
	return &demoServerImpl{demo, logger, tx}
}

func (s *demoServerImpl) SayHello(ctx context.Context, msg string) (res *domain.Demo, err error) {
	var (
		data *domain.Demo
	)
	data = s.repo.SayHello(ctx, msg)
	res = data
	return

}
