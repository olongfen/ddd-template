package repository

import (
	"context"
	"ddd-template/internal/domain"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

type demoRepo struct {
	data   *Data
	log    *zap.Logger
	bundle *i18n.Bundle
}

func (d *demoRepo) Get(ctx context.Context, demo *domain.Demo) error {
	if err := d.data.DB(ctx).Create(demo).Error; err != nil {
		return err
	}
	return nil
}

// NewDemoDependency
// #Description: new
// #param db *gorm.DB
// #return dependency.IDemoRepo
func NewDemoDependency(data *Data, bundle *i18n.Bundle, logger *zap.Logger) domain.IDemoRepo {
	return &demoRepo{data: data, bundle: bundle, log: logger}
}
