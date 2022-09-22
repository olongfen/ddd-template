package query

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
	"go.uber.org/zap"
)

type IStudentQueryService interface {
	GetStudent(ctx context.Context, uuid string) (ret *schema.StudentResp, err error)
}

type queryStudent struct {
	repo         domain.IStudentRepository
	classService domain.IClassDomainService
	logger       *zap.Logger
}

func (q queryStudent) GetStudent(ctx context.Context, uuid string) (ret *schema.StudentResp, err error) {
	student, err := q.repo.GetStudent(ctx, uuid)
	if err != nil {
		return nil, err
	}
	class, err := q.classService.GetClassDetail(ctx, student.ClassUuid())
	if err != nil {
		return nil, err
	}
	// return
	ret = schema.UnmarshalStudentFromEnt(student)
	ret.ClassName = class.Name()

	return
}

func NewQueryStudent(repo domain.IStudentRepository,
	classService domain.IClassDomainService,
	logger *zap.Logger) (ret IStudentQueryService) {
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
	u := new(queryStudent)
	u.repo = repo
	u.logger = logger
	u.classService = classService
	ret = u
	return
}
