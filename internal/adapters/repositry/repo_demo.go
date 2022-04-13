package repositry

import (
	"context"
	"ddd-template/internal/domain"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type demoDependency struct {
	db  *gorm.DB
	log *zap.Logger
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
// NewDemoDependencyImpl
// #Description: new
// #param db *gorm.DB
// #return dependency.IDemoRepo
func NewDemoDependencyImpl(db *gorm.DB, logger *zap.Logger) domain.IDemoRepo {
	return &demoDependency{db: db, log: logger}
}
