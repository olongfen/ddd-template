package domain

import (
	"ddd-template/pkg/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Class  entity
type Class struct {
	uuid      string
	createdAt utils.JSONTime
	updatedAt utils.JSONTime
	name      string
}

func (c Class) Uuid() string {
	return c.uuid
}

func (c Class) CreatedAt() utils.JSONTime {
	return c.createdAt
}

func (c Class) UpdatedAt() utils.JSONTime {
	return c.updatedAt
}

func (c Class) Name() string {
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
