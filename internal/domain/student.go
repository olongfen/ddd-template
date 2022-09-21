package domain

import (
	"context"
	"ddd-template/pkg/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Student entity
type Student struct {
	uuid      string
	createdAt utils.JSONTime
	updatedAt utils.JSONTime
	name      string
	stuNumber string
	classID   string
}

func (u Student) Uuid() string {
	return u.uuid
}

func (u Student) CreatedAt() utils.JSONTime {
	return u.createdAt
}

func (u Student) UpdatedAt() utils.JSONTime {
	return u.updatedAt
}

func (u Student) ClassID() string {
	return u.classID
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
	u.classID = classID
	u.stuNumber = stuNumber
	return
}

// IStudentRepository 用户表存储库
type IStudentRepository interface {
	AddStudent(ctx context.Context, stu *Student) (err error)
}
