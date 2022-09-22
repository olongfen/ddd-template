package mutation

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
	"go.uber.org/zap"
)

type classMutation struct {
	repo   domain.IClassRepository
	logger *zap.Logger
}

// AddClass 添加班级
func (c classMutation) AddClass(ctx context.Context, form *schema.ClassAddForm) (err error) {
	var (
		data *domain.Class
	)
	if data, err = domain.NewClass(form.Name); err != nil {
		return
	}
	return c.repo.AddClass(ctx, data)
}

type IClassMutationService interface {
	AddClass(ctx context.Context, form *schema.ClassAddForm) (err error)
}

func NewClassMutation(repo domain.IClassRepository,
	logger *zap.Logger) IClassMutationService {
	if repo == nil {
		zap.L().Fatal("empty repo")
	}
	if logger == nil {
		zap.L().Error("empty logger")
	}
	c := &classMutation{
		repo:   repo,
		logger: logger,
	}
	return c
}
