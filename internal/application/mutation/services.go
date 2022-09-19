package mutation

import "context"

type AddUserFrom struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IUserService interface {
	AddUser(ctx context.Context, form *AddUserFrom) (err error)
	DeleteUser(ctx context.Context, uuid string) (err error)
}

type userServiceHandler struct {
}
