package domain

import (
	"github.com/olongfen/toolkit/tools"
)

// Demo demo
type Demo struct {
	tools.Model
	Name string `gorm:"unique;index"`
}

type IDemoRepo interface {
	IRepository[Demo]
}
