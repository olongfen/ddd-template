package query

import (
	"context"
	"ddd-template/internal/application/transform"
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
	"go.uber.org/zap"
)

type IClassQueryService interface {
	QueryClasses(ctx context.Context, query *schema.ClassQueryReq) (ret schema.ClassRespList, pag *schema.Pagination, err error)
}

func (c queryClass) QueryClasses(ctx context.Context, query *schema.ClassQueryReq) (ret schema.ClassRespList, pag *schema.Pagination, err error) {
	var (
		data []*domain.Class
	)
	if len(query.Order) == 0 {
		query.Order = append(query.Order, "desc")
		query.Sort = append(query.Sort, "createdAt")
	}
	if data, pag, err = c.repo.QueryClasses(ctx, query); err != nil {
		return
	}
	for _, v := range data {
		ret = append(ret, transform.UnmarshalClassToSchema(v))
	}
	return
}

type queryClass struct {
	repo   domain.IClassRepository
	logger *zap.Logger
}

func NewQueryClass(repo domain.IClassRepository,
	logger *zap.Logger) (ret IClassQueryService) {
	if repo == nil {
		zap.L().Fatal("empty repository")
		return
	}
	if logger == nil {
		zap.L().Fatal("empty logger")
		return
	}
	u := new(queryClass)
	u.repo = repo
	u.logger = logger
	ret = u
	return
}
