package mutation

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
	"errors"
	"go.uber.org/zap"
)

type studentMutation struct {
	repo         domain.IStudentRepository
	classService domain.IClassDomainService
	logger       *zap.Logger
}

// AddStudent add
func (u studentMutation) AddStudent(ctx context.Context, form *schema.StudentAddForm) (err error) {
	var (
		stu *domain.Student
	)
	if _, err = u.classService.GetClassDetail(ctx, form.ClassUuid); err != nil {
		u.logger.Error("AddStudent", zap.Error(err))
		err = errors.New("class dose not exists")
		return
	}
	if stu, err = domain.NewStudent(form.Name, form.StuNumber, form.ClassUuid); err != nil {
		return
	}
	if err = u.repo.AddStudent(ctx, stu); err != nil {
		return
	}
	return
}

func NewUserMutation(repo domain.IStudentRepository,
	classService domain.IClassDomainService,
	logger *zap.Logger) (ret IStudentMutationService) {
	if repo == nil {
		zap.L().Fatal("empty repository")
		return
	}
	if classService == nil {
		zap.L().Fatal("empty class domain serve")
		return
	}
	if logger == nil {
		zap.L().Fatal("empty logger")
		return
	}
	u := new(studentMutation)
	u.repo = repo
	u.logger = logger
	u.classService = classService
	ret = u
	return
}

type IStudentMutationService interface {
	AddStudent(ctx context.Context, form *schema.StudentAddForm) (err error)
}