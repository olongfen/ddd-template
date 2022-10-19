package repository

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/pkg/utils"
	"go.uber.org/zap"
)

type Student struct {
	utils.Model
	StuNumber string `gorm:"uniqueIndex;comment:学号;size:10"`
	Name      string `gorm:"type:varchar(16);index;comment:名称"`
	ClassUuid string `gorm:"size:36;index;comment:班级id"`
}

type studentRepository struct {
	data *Data
}

// FindOne  get one
func (u studentRepository) FindOne(ctx context.Context, id int) (ret *domain.Student, err error) {
	var (
		data *Student
	)
	if err = u.data.DB(ctx).Model(&Student{}).Where("id = ?", id).First(&data).Error; err != nil {
		return
	}
	ret = u.unmarshal(data)
	return
}

// Find get page
func (u studentRepository) Find(ctx context.Context, o domain.OtherCond, fields ...domain.Field) (ret []*domain.Student,
	pagination *domain.Pagination, err error) {
	var (
		data []*Student
		db   = u.data.DB(ctx).Model(&Student{})
		opt  = newOption(o)
	)
	fieldsT(fields).process(db)
	if pagination, err = findPage(db, opt, &data); err != nil {
		return
	}
	for _, v := range data {
		ret = append(ret, u.unmarshal(v))
	}
	return
}

func (u studentRepository) unmarshal(data *Student) *domain.Student {
	return domain.UnmarshalStudentFromDatabase(data.ID, data.Uuid, data.CreatedAt, data.UpdatedAt, data.Name, data.StuNumber, data.ClassUuid)
}

func (u studentRepository) marshal(in *domain.Student) *Student {
	stu := &Student{
		Model: utils.Model{
			Uuid: in.Uuid(),
		},
		Name:      in.Name(),
		StuNumber: in.StuNumber(),
		ClassUuid: in.ClassUuid(),
	}
	return stu
}

// Create 往数据库写入user记录
func (u studentRepository) Create(ctx context.Context, stu *domain.Student) (err error) {
	if err = u.data.DB(ctx).Model(&Student{}).Create(u.marshal(stu)).Error; err != nil {
		return
	}
	return
}

// Update update
func (u studentRepository) Update(ctx context.Context, id int, stu *domain.Student) (err error) {
	if err = u.data.DB(ctx).Model(&Student{}).Where("id = ?", id).Updates(u.marshal(stu)).Error; err != nil {
		return
	}
	return
}

// Delete del
func (u studentRepository) Delete(ctx context.Context, id int) (err error) {
	if err = u.data.DB(ctx).Where("id = ?", id).Delete(&Student{}).Error; err != nil {
		return
	}
	return
}

// NewStudentRepository new user repository
func NewStudentRepository(data *Data) (ret domain.IStudentRepository) {
	if data == nil {
		zap.L().Fatal("empty data")
		return
	}
	stu := &studentRepository{
		data: data,
	}
	ret = stu
	return
}
