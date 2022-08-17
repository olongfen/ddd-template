package repository

import (
	"context"
	"ddd-template/internal/domain"
	"go.uber.org/zap"
)

type demoRepo struct {
	data *Data
	log  *zap.Logger
}

func (d *demoRepo) Get(ctx context.Context, demo *domain.Demo) error {
	if err := d.data.DB(ctx).First(demo).Error; err != nil {
		return err
	}
	return nil
}

// NewDemoDependency
// #Description: new
// #param db *gorm.DB
// #return dependency.IDemoRepo
func NewDemoDependency(data *Data, logger *zap.Logger) domain.IDemoRepo {
	return &demoRepo{data: data, log: logger}
}
