package schema

import (
	"time"
)

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
	CreatedAt time.Time `json:"createdAt"`
	// 更新时间
	UpdatedAt time.Time `json:"updatedAt"`
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
