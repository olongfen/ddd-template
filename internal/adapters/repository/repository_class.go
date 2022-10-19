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
	data *Data
}

// Find query
func (c classRepository) Find(ctx context.Context, o domain.OtherCond, fields ...domain.Field) (ret []*domain.Class, pag *domain.Pagination, err error) {
	var (
		data []*Class
		db   = c.data.DB(ctx).Model(&Class{})
		opt  = newOption(o)
	)

	fieldsT(fields).process(db)

	if pag, err = findPage(db, opt, &data); err != nil {
		return
	}
	for _, v := range data {
		ret = append(ret, c.unmarshal(v))
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
	ret = cls
	return
}

// FindByUuid find by uuid
func (c classRepository) FindByUuid(ctx context.Context, uid string) (ret *domain.Class, err error) {
	var (
		m = new(Class)
	)
	if err = c.data.DB(ctx).Model(&Class{}).Where("uuid = ?", uid).First(m).Error; err != nil {
		return
	}
	ret = c.unmarshal(m)
	return
}

// Create one
func (c classRepository) Create(ctx context.Context, cls *domain.Class) (err error) {
	return c.data.DB(ctx).Create(c.marshal(cls)).Error
}

// Update update
func (c classRepository) Update(ctx context.Context, id int, ent *domain.Class) (err error) {
	return c.data.DB(ctx).Model(&Class{}).Where("id = ?", id).Updates(c.marshal(ent)).Error
}

func (c classRepository) FindOne(ctx context.Context, id int) (ret *domain.Class, err error) {
	var (
		data = new(Class)
	)
	if err = c.data.DB(ctx).Where("id = ?", id).First(data).Error; err != nil {
		return
	}
	ret = c.unmarshal(data)
	return
}

func (c classRepository) Delete(ctx context.Context, id int) (err error) {
	return c.data.DB(ctx).Where("id = ?", id).Delete(&Class{}).Error

}

func (c classRepository) unmarshal(data *Class) *domain.Class {
	return domain.UnmarshalClassFromDatabase(data.ID, data.Uuid, data.CreatedAt, data.UpdatedAt, data.Name)
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
