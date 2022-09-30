package utils

import "gorm.io/gorm"

type Model struct {
	ID        uint           `gorm:"primarykey"`
	Uuid      string         `gorm:"size:36;uniqueIndex;comment:唯一uuid"`
	CreatedAt JSONTime       `gorm:"comment:创建时间"`
	UpdatedAt JSONTime       `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
}
