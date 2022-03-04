package entities

import "gorm.io/gorm"

type Demo struct {
	gorm.Model
	Message string
}
