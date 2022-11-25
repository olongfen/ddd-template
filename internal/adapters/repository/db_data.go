package repository

import (
	"context"
	"gorm.io/gorm"
)

type DBData interface {
	DB(ctx context.Context) *gorm.DB
	Close() error
}
