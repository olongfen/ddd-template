package domain

import (
	"context"
	"go.uber.org/zap"
)

type ClassDomainService struct {
	repo   IClassRepository
	logger *zap.Logger
}

func NewClassDomainService(repo IClassRepository, logger *zap.Logger) IClassDomainService {
	if repo == nil {
		zap.L().Fatal("empty class repository")
	}
	if logger == nil {
		zap.L().Fatal("empty logger")
	}
	ret := new(ClassDomainService)
	ret.repo = repo
	ret.logger = logger
	return ret
}

func (c ClassDomainService) GetClassDetail(ctx context.Context, uid string) (ret *Class, err error) {
	return c.repo.FindByUuid(ctx, uid)
}
