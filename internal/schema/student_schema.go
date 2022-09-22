package schema

type StudentAddForm struct {
	// 姓名
	Name string `json:"name" validate:"required"`
	// 班级uuid
	ClassUuid string `json:"classUuid" validate:"required,uuid4"`
	// 学号
	StuNumber string `json:"stuNumber" validate:"required"`
}
