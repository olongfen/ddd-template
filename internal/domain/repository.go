package domain

import "context"

type IRepository[T any] interface {
	Create(ctx context.Context, ent *T) (err error)
	FindOne(ctx context.Context, id int) (ret *T, err error)
	FindOneBy(ctx context.Context, field ...Field) (ret *T, err error)
	Find(ctx context.Context, o OtherCond, fields ...Field) (ret []*T,
		pagination *Pagination, err error)
	Update(ctx context.Context, id int, stu *T) (err error)
	Delete(ctx context.Context, id int) (err error)
}
