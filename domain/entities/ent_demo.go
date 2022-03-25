package entities

import (
	"ddd-template/common/schema"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type Demo struct {
	gorm.Model
	Message string
}

func (obj Demo) ToSchema() (res *schema.DemoInfo, err error) {
	var (
		to = new(schema.DemoInfo)
	)
	if err = copier.Copy(to, obj); err != nil {
		return
	}
	return to, nil
}

func (obj *Demo) SetID(param uint) *Demo {
	obj.ID = param
	return obj
}

func (obj *Demo) GetID() uint {
	return obj.ID
}

func (obj *Demo) SetCreatedAt(param time.Time) *Demo {
	obj.CreatedAt = param
	return obj
}

func (obj *Demo) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

func (obj *Demo) SetUpdatedAt(param time.Time) *Demo {
	obj.UpdatedAt = param
	return obj
}

func (obj *Demo) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj *Demo) SetMessage(param string) *Demo {
	obj.Message = param
	return obj
}

func (obj *Demo) GetMessage() string {
	return obj.Message
}
