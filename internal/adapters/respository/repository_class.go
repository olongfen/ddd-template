package respository

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/pkg/utils"
	"go.uber.org/zap"
)

type Class struct {
	utils.Model
	Name string `gorm:"size:24;comment:班级名称"`
}

type classRepository struct {
	data *Data
}

// NewClassRepository new class repository
func NewClassRepository(data *Data) (ret domain.IClassRepository) {
	if data == nil {
		zap.L().Fatal("empty data")
	}
	cls := new(classRepository)
	cls.data = data
	ret = cls
	return
}

func (c classRepository) GetClassWithUuid(ctx context.Context, uid string) (ret *domain.Class, err error) {
	var (
		m = new(Class)
	)
	if err = c.data.DB(ctx).Model(&Class{}).Where("uuid = ?", uid).First(m).Error; err != nil {
		return
	}
	ret = domain.UnmarshalClassFromDatabase(m.Uuid, m.CreatedAt, m.UpdatedAt, m.Name)
	return
}

func (c classRepository) AddClass(ctx context.Context, cls *domain.Class) (err error) {
	return c.data.DB(ctx).Create(c.marshal(cls)).Error
}

func (c classRepository) marshal(in *domain.Class) *Class {
	class := &Class{
		Model: utils.Model{
			Uuid: in.Uuid(),
		},
		Name: in.Name(),
	}
	return class
}
