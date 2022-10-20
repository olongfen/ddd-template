package transform

import (
	"ddd-template/internal/application/schema"
	"ddd-template/internal/domain"
)

func UnmarshalClassToSchema(c *domain.Class) *schema.ClassResp {
	return &schema.ClassResp{
		Name:      c.Name,
		ClassUuid: c.Uuid,
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func UnmarshalStudentToSchema(ent *domain.Student) *schema.StudentResp {
	return &schema.StudentResp{
		Uuid:      ent.Uuid,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
		Name:      ent.Name,
		StuNumber: ent.StuNumber,
		ClassUuid: ent.ClassUuid,
		ID:        ent.ID,
		ClassName: "",
	}
}
