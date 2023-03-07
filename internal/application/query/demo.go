package query

import (
	"context"
	"ddd-template/internal/application/query/query_iface"
	"ddd-template/internal/domain"
)

// demo demo
type demo struct {
	repo domain.IDemoRepo
}

// Hello  world
func (d *demo) Hello(ctx context.Context) string {
	return "hello world"
}

// NewDemo new
func NewDemo(repo domain.IDemoRepo) query_iface.IDemoService {
	return &demo{
		repo: repo,
	}
}
