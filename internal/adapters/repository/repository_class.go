package repository

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/pkg/utils"
	"go.uber.org/zap"
)

type Class struct {
	utils.Model
	Name string `gorm:"size:24;uniqueIndex;comment:班级名称"`
}

type classRepository struct {
	repository[domain.Class]
	data *Data
}

func (c classRepository) FindByUuid(ctx context.Context, uid string) (ret *domain.Class, err error) {
	if err = c.data.db.WithContext(ctx).Model(&domain.Class{}).Where("uuid = ?", uid).First(&ret).Error; err != nil {
		return
	}
	return
}

// NewClassRepository new class repository
func NewClassRepository(data *Data) (ret domain.IClassRepository) {
	if data == nil {
		zap.L().Fatal("empty data")
	}
	cls := new(classRepository)
	cls.data = data
	cls.repository = repository[domain.Class]{
		data: data,
	}
	ret = cls
	return
}
