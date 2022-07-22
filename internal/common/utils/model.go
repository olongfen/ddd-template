package utils

import "gorm.io/gorm"

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt JSONTime
	UpdatedAt JSONTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
