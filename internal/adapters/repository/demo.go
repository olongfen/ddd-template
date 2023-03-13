package repository

import (
	"ddd-template/internal/domain"
	"github.com/olongfen/toolkit/db_data"
)

// demo demo
type demo struct {
	repository[domain.Demo]
}

// NewDemo new demo repo
func NewDemo(db db_data.DBData) domain.IDemoRepo {
	return &demo{repository[domain.Demo]{data: db}}
}
