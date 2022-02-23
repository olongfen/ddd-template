package repositry

import (
	"context"
	"ddd-template/domain/dependency"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type demoDependencyImpl struct {
	db  *gorm.DB
	log *zap.Logger
}

//
// SayHello
// #Description: say hello
// #receiver d *demoDependencyImpl
// #param ctx context.Context
// #param msg string
// #return string
func (d *demoDependencyImpl) SayHello(ctx context.Context, msg string) string {
	return msg + " hello"
}

//
// NewDemoDependencyImpl
// #Description: new
// #param db *gorm.DB
// #return dependency.DemoInterface
func NewDemoDependencyImpl(db *gorm.DB, logger *zap.Logger) dependency.DemoInterface {
	return &demoDependencyImpl{db: db, log: logger}
}
