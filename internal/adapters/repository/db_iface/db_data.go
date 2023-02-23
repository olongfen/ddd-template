package db_iface

import (
	"context"
	"gorm.io/gorm"
)

type DBData interface {
	DB(ctx context.Context) *gorm.DB
	ExecTx(ctx context.Context, fc func(context.Context) error) error
	Close() error
}

type ITransaction interface {
	ExecTx(ctx context.Context, fc func(ctx context.Context) error) error
}
