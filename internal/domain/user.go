package domain

import (
	"context"
	"ddd-template/pkg/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// User user
type User struct {
	uuid      string
	createdAt utils.JSONTime
	updatedAt utils.JSONTime
	username  string
	password  string
}

func (u User) UUID() string {
	return u.uuid
}

func (u User) Username() string {
	return u.username
}

func (u User) Password() string {
	return u.password
}

// NewUser new user
func NewUser(username string, password string) (u *User, err error) {
	if len(username) == 0 {
		err = errors.New("empty username")
		return
	}
	if len(password) == 0 {
		err = errors.New("empty password")
		return
	}
	u = &User{}
	u.uuid = uuid.NewV4().String()
	u.username = username
	u.password = password
	return
}

// IUserRepository 用户表存储库
type IUserRepository interface {
	AddUser(ctx context.Context, user *User) (err error)
}
