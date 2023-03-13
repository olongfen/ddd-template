package query

import (
	"context"
	"ddd-template/internal/domain"
)

// demo demo
type demo struct {
	repo domain.IDemoRepo
}

// Hello  world
func (d *demo) Hello(ctx context.Context, msg string) string {
	return "hello, " + msg
}

// NewDemo new
func NewDemo(repo domain.IDemoRepo) IDemoService {
	return &demo{
		repo: repo,
	}
}
