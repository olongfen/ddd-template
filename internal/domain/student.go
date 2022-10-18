package domain

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Student entity
type Student struct {
	id        uint
	uuid      string
	createdAt time.Time
	updatedAt time.Time
	name      string
	stuNumber string
	classUuid string
}

func (u Student) Id() uint {
	return u.id
}

func UnmarshalStudentFromDatabase(
	id uint,
	uuid string,
	createdAt time.Time,
	updatedAt time.Time,
	name string,
	stuNumber string,
	classUuid string) *Student {
	return &Student{
		id:        id,
		uuid:      uuid,
		createdAt: createdAt,
		updatedAt: updatedAt,
		name:      name,
		stuNumber: stuNumber,
		classUuid: classUuid,
	}
}

func (u Student) Uuid() string {
	return u.uuid
}

func (u Student) CreatedAt() time.Time {
	return u.createdAt
}

func (u Student) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u Student) ClassUuid() string {
	return u.classUuid
}

func (u Student) Name() string {
	return u.name
}

func (u Student) StuNumber() string {
	return u.stuNumber
}

// NewStudent new student
func NewStudent(name string, stuNumber string, classID string) (u *Student, err error) {
	if len(name) == 0 {
		err = errors.New("empty name")
		return
	}
	if len(stuNumber) == 0 {
		err = errors.New("empty stuNumber")
		return
	}
	if len(classID) == 0 {
		err = errors.New("empty class")
		return
	}
	u = &Student{}
	u.uuid = uuid.NewV4().String()
	u.name = name
	u.classUuid = classID
	u.stuNumber = stuNumber
	return
}

// IStudentRepository 用户表存储库
type IStudentRepository interface {
	AddStudent(ctx context.Context, stu *Student) (err error)
	GetStudent(ctx context.Context, id int) (ret *Student, err error)
	FindStudent(ctx context.Context, o OtherCond, fields ...Field) (ret []*Student,
		pagination *Pagination, err error)
	UpStudent(ctx context.Context, id uint, stu *Student) (err error)
}

// UnmarshalStudentFromSchemaUpForm 把更新结构体赋值给实体
func UnmarshalStudentFromSchemaUpForm(name string, classUuid string) *Student {
	return &Student{
		name:      name,
		classUuid: classUuid,
	}
}
