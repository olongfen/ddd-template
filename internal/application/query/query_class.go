package query

import (
	"context"
	"ddd-template/internal/application/schema"
	"ddd-template/internal/application/transform"
	"ddd-template/internal/domain"
	"ddd-template/pkg/utils"
	"go.uber.org/zap"
)

type IClassQueryService interface {
	QueryClasses(ctx context.Context, query *schema.ClassQueryReq) (ret schema.ClassRespList, pag *schema.Pagination, err error)
	GetClass(ctx context.Context, id int) (ret *schema.ClassResp, err error)
}

func (c queryClass) GetClass(ctx context.Context, id int) (ret *schema.ClassResp, err error) {
	var (
		data *domain.Class
	)
	if data, err = c.repo.FindOne(ctx, id); err != nil {
		return
	}
	ret = transform.UnmarshalClassToSchema(data)
	return
}

func (c queryClass) QueryClasses(ctx context.Context, query *schema.ClassQueryReq) (ret schema.ClassRespList, pag *schema.Pagination, err error) {
	var (
		data []*domain.Class
		dPag *domain.Pagination
	)
	if len(query.Order) == 0 {
		query.Order = append(query.Order, "desc")
		query.Sort = append(query.Sort, "createdAt")
	}
	if data, dPag, err = c.repo.Find(ctx, domain.OtherCond{
		PageSize:    query.PageSize,
		CurrentPage: query.CurrentPage,
		Sort:        query.Sort,
		Order:       query.Order,
	}); err != nil {
		return
	}
	for _, v := range data {
		ret = append(ret, transform.UnmarshalClassToSchema(v))
	}
	utils.MustDecode(dPag, &pag)
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
