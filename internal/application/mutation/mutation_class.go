package mutation

import (
	"context"
	"ddd-template/internal/application/schema"
	"ddd-template/internal/domain"
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

// UpClass update
func (c classMutation) UpClass(ctx context.Context, form *schema.ClassUpForm) (err error) {
	if err = c.repo.UpClass(ctx, form.Id, domain.UnmarshalClassFromSchemaUpForm(form.Name)); err != nil {
		return
	}
	return
}

type IClassMutationService interface {
	AddClass(ctx context.Context, form *schema.ClassAddForm) (err error)
	UpClass(ctx context.Context, form *schema.ClassUpForm) (err error)
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
