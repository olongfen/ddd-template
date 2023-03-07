package repository

import (
	"ddd-template/internal/adapters/repository/db_iface"
	"ddd-template/internal/domain"
)

// demo demo
type demo struct {
	repository[domain.Demo]
}

// NewDemo new demo repo
func NewDemo(db db_iface.DBData) domain.IDemoRepo {
	return &demo{repository[domain.Demo]{data: db}}
}
