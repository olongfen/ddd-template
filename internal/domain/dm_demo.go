package domain

import (
	"context"
	"ddd-template/internal/common/utils"
)

type IDemoRepo interface {
	Get(ctx context.Context, demo *Demo) error
}

type IDemoUseCase interface {
	Get(ctx context.Context, id uint) (demo *Demo, err error)
}

type Demo struct {
	utils.Model
	Message string
}
