package mutation

import (
	"ddd-template/internal/application/mutation/mutat_iface"
	"ddd-template/internal/domain"
)

// demo demo
type demo struct {
	repo domain.IDemoRepo
}

// NewDemo new
func NewDemo(repo domain.IDemoRepo) mutat_iface.IDemoService {
	return &demo{repo: repo}
}
