package domain

import (
	"context"
	"ddd-template/internal/common/utils"
)

type IDemoRepo interface {
	SayHello(ctx context.Context, msg string) *Demo
}

type IDemoUsecase interface {
	SayHello(ctx context.Context, msg string) (*Demo, error)
}

type Demo struct {
	utils.Model
	Message string
}
