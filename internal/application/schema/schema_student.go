package schema

import (
	"ddd-template/pkg/utils"
)

// StudentAddForm 添加一条学生记录
type StudentAddForm struct {
	// 姓名
	Name string `json:"name" validate:"required"`
	// 班级uuid
	ClassUuid string `json:"classUuid" validate:"required,uuid4"`
	// 学号
	StuNumber string `json:"stuNumber" validate:"required,min=1,max=10"`
}

// StudentResp 返回结构体
type StudentResp struct {
	// ID
	ID uint `json:"id"`
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

// StudentsResp 列表返回体
type StudentsResp []*StudentResp

type StudentsQueryResp struct {
	List       StudentsResp `json:"list"`
	Pagination *Pagination  `json:"pagination"`
}

// StudentsQuery 分页查询
type StudentsQuery struct {
	Name string `json:"name" query:"name"`
	QueryOptions
}

// StudentUpForm 更新学生信息
type StudentUpForm struct {
	// Name 姓名
	Name string `json:"name"`
	// 班级uuid
	ClassUuid string `json:"classUuid" validate:"uuid4"`
}
