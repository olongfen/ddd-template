package schema

import (
	"ddd-template/pkg/utils"
)

type StudentAddForm struct {
	// 姓名
	Name string `json:"name" validate:"required"`
	// 班级uuid
	ClassUuid string `json:"classUuid" validate:"required,uuid4"`
	// 学号
	StuNumber string `json:"stuNumber" validate:"required,min=1,max=10"`
}

type StudentResp struct {
	// Uuid
	Uuid string `json:"uuid"`
	// 创建时间
	CreatedAt utils.JSONTime `json:"createdAt"`
	// 更新时间
	UpdatedAt utils.JSONTime `json:"updatedAt"`
	// 学生姓名
	Name string `json:"name"`
	// 学号
	StuNumber string `json:"stuNumber"`
	// 班级uuid
	ClassUuid string `json:"classUuid"`
	// 班级名称
	ClassName string `json:"className"`
}

type StudentsResp []*StudentResp

type StudentsQuery struct {
	QueryOptions
}
