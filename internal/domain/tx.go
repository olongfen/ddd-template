package domain

import "context"

type ITransaction interface {
	ExecTx(ctx context.Context, fc func(ctx context.Context) error) error
}
