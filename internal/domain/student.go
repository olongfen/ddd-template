package domain

import (
	"ddd-template/pkg/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Student entity
type Student struct {
	utils.Model
	StuNumber string `gorm:"uniqueIndex;comment:学号;size:10"`
	Name      string `gorm:"type:varchar(16);index;comment:名称"`
	ClassUuid string `gorm:"size:36;index;comment:班级id"`
}

// NewStudent new student
func NewStudent(name string, stuNumber string, classID string) (u *Student, err error) {
	if len(name) == 0 {
		err = errors.New("empty Name")
		return
	}
	if len(stuNumber) == 0 {
		err = errors.New("empty StuNumber")
		return
	}
	if len(classID) == 0 {
		err = errors.New("empty class")
		return
	}
	u = &Student{}
	u.Uuid = uuid.NewV4().String()
	u.Name = name
	u.ClassUuid = classID
	u.StuNumber = stuNumber
	return
}

// IStudentRepository 用户表存储库
type IStudentRepository interface {
	IRepository[Student]
}

// UnmarshalStudentFromSchemaUpForm 把更新结构体赋值给实体
func UnmarshalStudentFromSchemaUpForm(name string, classUuid string) *Student {
	return &Student{
		Name:      name,
		ClassUuid: classUuid,
	}
}
