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

func NewDemoServer(demo domain.IDemoRepo, tx app.ITransaction, logger *zap.Logger) domain.IDemoUseCase {
	return &demoServerImpl{demo, logger, tx}
}

func (s *demoServerImpl) Get(ctx context.Context, id uint) (ret *domain.Demo, err error) {
	ret = new(domain.Demo)
	ret.ID = id
	if err = s.repo.Get(ctx, ret); err != nil {
		return
	}
	return

}
