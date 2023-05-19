package command

import (
	"ddd-template/internal/domain"
)

// demo demo
type demo struct {
	repo domain.IDemoRepo
}

// NewDemo new
func NewDemo(repo domain.IDemoRepo) IDemoService {
	return &demo{repo: repo}
}
