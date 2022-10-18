package schema

import "ddd-template/pkg/utils"

type ClassAddForm struct {
	// 班级名称
	Name string `json:"name" validate:"required"`
}

type ClassResp struct {
	// 班级名称
	Name string `json:"name"`
	// class uuid
	ClassUuid string `json:"classUuid"`
	// id
	ID uint `json:"id"`
	// 创建时间
	CreatedAt utils.JSONTime `json:"createdAt"`
	// 更新时间
	UpdatedAt utils.JSONTime `json:"updatedAt"`
}

type ClassRespList []*ClassResp

type ClassQueryReq struct {
	QueryOptions
}

type ClassQueryResp struct {
	List       ClassRespList `json:"list"`
	Pagination *Pagination   `json:"pagination"`
}

// ClassUpForm update form
type ClassUpForm struct {
	// path param
	Id int `json:"-" params:"id" validate:"required,min=1"`
	// update body
	Name string `json:"name"`
}
