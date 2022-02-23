package service

import (
	"context"
	"ddd-template/domain/dependency"
	"go.uber.org/zap"
)

type DemoService struct {
	repo dependency.DemoInterface
	log  *zap.Logger
}

//
// NewDemoService
// #Description: new
// #param repo dependency.DemoInterface
// #return *DemoService
func NewDemoService(repo dependency.DemoInterface, logger *zap.Logger) *DemoService {
	return &DemoService{repo: repo, log: logger}
}

//
// SayHello
// #Description: say hello
// #receiver d *DemoService
// #param ctx context.Context
// #param msg string
// #return string
func (d *DemoService) SayHello(ctx context.Context, msg string) string {
	return d.repo.SayHello(ctx, msg)
}
