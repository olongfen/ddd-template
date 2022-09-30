package query

import (
	"context"
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
	"go.uber.org/zap"
)

type IStudentQueryService interface {
	GetStudent(ctx context.Context, uuid string) (ret *schema.StudentResp, err error)
	QueryStudents(ctx context.Context, query schema.StudentsQuery) (ret schema.StudentsResp,
		pagination *schema.Pagination, err error)
}

type queryStudent struct {
	repo         domain.IStudentRepository
	classService domain.IClassDomainService
	logger       *zap.Logger
}

// GetStudent get
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
	ret = domain.UnmarshalStudentToSchema(student)
	ret.ClassName = class.Name()

	return
}

// QueryStudents query students by page
func (q queryStudent) QueryStudents(ctx context.Context, query schema.StudentsQuery) (ret schema.StudentsResp,
	pagination *schema.Pagination, err error) {
	var (
		data []*domain.Student
	)
	if data, pagination, err = q.repo.QueryStudents(ctx, query); err != nil {
		return
	}

	for _, v := range data {
		class, _err := q.classService.GetClassDetail(ctx, v.ClassUuid())
		if _err != nil {
			err = _err
			return
		}
		_data := domain.UnmarshalStudentToSchema(v)
		_data.ClassName = class.Name()
		ret = append(ret, _data)
	}
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
