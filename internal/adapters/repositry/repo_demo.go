package repositry

import (
	"context"
	"ddd-template/internal/domain"
	"go.uber.org/zap"
)

type demoDependency struct {
	data *Data
	log  *zap.Logger
}

//
// SayHello
// #Description: say hello
// #receiver d *demoDependency
// #param ctx context.Context
// #param msg string
// #return string
func (d *demoDependency) SayHello(ctx context.Context, msg string) *domain.Demo {
	res := new(domain.Demo)
	res.Message = msg + " hello"
	return res
}

//
// NewDemoDependency
// #Description: new
// #param db *gorm.DB
// #return dependency.IDemoRepo
func NewDemoDependency(data *Data, logger *zap.Logger) domain.IDemoRepo {
	return &demoDependency{data: data, log: logger}
}
