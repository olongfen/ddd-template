package schema

type ClassAddForm struct {
	// 班级名称
	Name string `json:"name" validate:"required"`
}
