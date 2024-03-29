package domain

import "context"

// IRepository 模型泛型存储库接口
type IRepository[T any] interface {
	Create(ctx context.Context, ent *T) (err error)
	FindOne(ctx context.Context, id int) (ret *T, err error)
	FindOneBy(ctx context.Context, field ...Field) (ret *T, err error)
	Find(ctx context.Context, o OtherCond, fields ...Field) (ret []*T,
		pagination *Pagination, err error)
	Update(ctx context.Context, id int, stu *T) (err error)
	Delete(ctx context.Context, id int) (err error)
	DeleteBy(ctx context.Context, fields ...Field) (err error)
	Count(ctx context.Context, fields ...Field) (count int64, err error)
}
