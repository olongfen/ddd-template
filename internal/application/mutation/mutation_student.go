package mutation

import (
	"context"
	"ddd-template/internal/application/schema"
	"ddd-template/internal/domain"
	"errors"
	"go.uber.org/zap"
)

type studentMutation struct {
	repo         domain.IStudentRepository
	classService domain.IClassDomainService
	logger       *zap.Logger
}

// UpStudent update
func (u studentMutation) UpStudent(ctx context.Context, id uint, form *schema.StudentUpForm) (err error) {

	// 判断班级是否存在
	if _, err = u.classService.GetClassDetail(ctx, form.ClassUuid); err != nil {
		return
	}
	return u.repo.UpStudent(ctx, id, domain.UnmarshalStudentFromSchemaUpForm(form.Name, form.ClassUuid))
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
	UpStudent(ctx context.Context, id uint, form *schema.StudentUpForm) (err error)
}
