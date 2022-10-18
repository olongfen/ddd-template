package domain

import (
	"context"
	"ddd-template/pkg/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Class  entity
type Class struct {
	id        uint
	uuid      string
	createdAt utils.JSONTime
	updatedAt utils.JSONTime
	name      string
}

func (c *Class) Id() uint {
	return c.id
}

func UnmarshalClassFromDatabase(
	id uint,
	uid string, createdAt utils.JSONTime,
	updatedAt utils.JSONTime,
	name string,
) *Class {
	return &Class{
		id:        id,
		uuid:      uid,
		createdAt: createdAt,
		updatedAt: updatedAt,
		name:      name,
	}
}

func (c *Class) Uuid() string {
	return c.uuid
}

func (c *Class) CreatedAt() utils.JSONTime {
	return c.createdAt
}

func (c *Class) UpdatedAt() utils.JSONTime {
	return c.updatedAt
}

func (c *Class) Name() string {
	return c.name
}

// NewClass new
func NewClass(name string) (c *Class, err error) {
	if name == "" {
		err = errors.New("empty name")
		return
	}
	c = new(Class)
	c.name = name
	c.uuid = uuid.NewV4().String()
	return
}

// UnmarshalClassFromSchemaUpForm 转换
func UnmarshalClassFromSchemaUpForm(name string) *Class {
	return &Class{
		name: name,
	}
}

// IClassRepository class repository
type IClassRepository interface {
	GetClassWithUuid(ctx context.Context, uid string) (ret *Class, err error)
	AddClass(ctx context.Context, c *Class) (err error)
	UpClass(ctx context.Context, id int, ent *Class) (err error)
	FindClass(ctx context.Context, o OtherCond, fields ...Field) (ret []*Class, pag *Pagination, err error)
}

// IClassDomainService domain serve
type IClassDomainService interface {
	GetClassDetail(ctx context.Context, uid string) (ret *Class, err error)
}
