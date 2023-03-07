package domain

import "github.com/olongfen/toolkit/utils"

// Demo demo
type Demo struct {
	utils.Model
	Name string
}

type IDemoRepo interface {
	IRepository[Demo]
}
