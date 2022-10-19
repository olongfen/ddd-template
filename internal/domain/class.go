package domain

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Class  entity
type Class struct {
	id        uint
	uuid      string
	createdAt time.Time
	updatedAt time.Time
	name      string
}

func (c *Class) Id() uint {
	return c.id
}

func UnmarshalClassFromDatabase(
	id uint,
	uid string, createdAt time.Time,
	updatedAt time.Time,
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

func (c *Class) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Class) UpdatedAt() time.Time {
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
	FindByUuid(ctx context.Context, uid string) (ret *Class, err error)
	Create(ctx context.Context, c *Class) (err error)
	Update(ctx context.Context, id int, ent *Class) (err error)
	Find(ctx context.Context, o OtherCond, fields ...Field) (ret []*Class, pag *Pagination, err error)
	FindOne(ctx context.Context, id int) (ret *Class, err error)
	Delete(ctx context.Context, id int) (err error)
}

// IClassDomainService domain serve
type IClassDomainService interface {
	GetClassDetail(ctx context.Context, uid string) (ret *Class, err error)
}
