package domain

import "context"

// ITransaction 事务
type ITransaction interface {
	ExecTx(ctx context.Context, fc func(c context.Context) error) error
}
