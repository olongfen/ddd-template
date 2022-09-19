package utils

import "gorm.io/gorm"

type Model struct {
	UUID      string         `gorm:"size:36;primarykey;comment:唯一uuid"`
	CreatedAt JSONTime       `gorm:"comment:创建时间"`
	UpdatedAt JSONTime       `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
}
