package repository

import (
	"ddd-template/internal/domain"
	"go.uber.org/zap"
)

type studentRepository struct {
	repository[domain.Student]
	data *Data
}

// NewStudentRepository new user repository
func NewStudentRepository(data *Data) (ret domain.IStudentRepository) {
	if data == nil {
		zap.L().Fatal("empty data")
		return
	}
	stu := &studentRepository{
		data: data,
		repository: repository[domain.Student]{
			data: data,
		},
	}
	ret = stu
	return
}
