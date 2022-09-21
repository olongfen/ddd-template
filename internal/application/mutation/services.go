package mutation

import (
	"context"
	"ddd-template/internal/domain"
	"go.uber.org/zap"
)

type AddUserFrom struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IUserService interface {
	AddUser(ctx context.Context, form *AddUserFrom) (err error)
	DeleteUser(ctx context.Context, uuid string) (err error)
}

type userServiceHandler struct {
	repo   domain.IStudentRepository
	tx     domain.ITransaction
	logger *zap.Logger
}

func (u userServiceHandler) AddUser(ctx context.Context, form *AddUserFrom) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceHandler) DeleteUser(ctx context.Context, uuid string) (err error) {
	//TODO implement me
	panic("implement me")
}

// NewUserServiceHandler
func NewUserServiceHandler(repo domain.IStudentRepository,
	tx domain.ITransaction) IUserService {
	svc := &userServiceHandler{
		repo: repo,
		tx:   tx,
	}
	return svc
}
