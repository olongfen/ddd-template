package respository

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

type userRepository struct {
	data *Data
}

func (u userRepository) marshal(in *domain.Student) *Student {
	stu := &Student{
		Model: utils.Model{
			Uuid: in.Uuid(),
		},
		Name:      in.Name(),
		StuNumber: in.StuNumber(),
	}
	return stu
}

// AddStudent 往数据库写入user记录
func (u userRepository) AddStudent(ctx context.Context, stu *domain.Student) (err error) {
	if err = u.data.DB(ctx).Model(&Student{}).Create(u.marshal(stu)).Error; err != nil {
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
	stu := &userRepository{
		data: data,
	}
	ret = stu
	return
}
