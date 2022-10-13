package transform

import (
	"ddd-template/internal/domain"
	"ddd-template/internal/schema"
)

func UnmarshalClassToSchema(c *domain.Class) *schema.ClassResp {
	return &schema.ClassResp{
		Name:      c.Name(),
		ClassUuid: c.Uuid(),
		ID:        c.Id(),
		CreatedAt: c.CreatedAt(),
		UpdatedAt: c.UpdatedAt(),
	}
}

func UnmarshalStudentToSchema(ent *domain.Student) *schema.StudentResp {
	return &schema.StudentResp{
		Uuid:      ent.Uuid(),
		CreatedAt: ent.CreatedAt(),
		UpdatedAt: ent.UpdatedAt(),
		Name:      ent.Name(),
		StuNumber: ent.StuNumber(),
		ClassUuid: ent.ClassUuid(),
		ID:        ent.Id(),
		ClassName: "",
	}
}
