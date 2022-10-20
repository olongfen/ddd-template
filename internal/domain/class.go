package domain

import (
	"context"
	"ddd-template/pkg/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Class  entity
type Class struct {
	utils.Model
	Name string
}

// NewClass new
func NewClass(name string) (c *Class, err error) {
	if name == "" {
		err = errors.New("empty Name")
		return
	}
	c = new(Class)
	c.Name = name
	c.Uuid = uuid.NewV4().String()
	return
}

// UnmarshalClassFromSchemaUpForm 转换
func UnmarshalClassFromSchemaUpForm(name string) *Class {
	return &Class{
		Name: name,
	}
}

// IClassRepository class repository
type IClassRepository interface {
	IRepository[Class]
	FindByUuid(ctx context.Context, uid string) (ret *Class, err error)
}

// IClassDomainService domain serve
type IClassDomainService interface {
	GetClassDetail(ctx context.Context, uid string) (ret *Class, err error)
}
