package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type ITransaction interface {
	ExecTx(ctx context.Context, fc func(ctx context.Context) error) error
}

type IHandler interface {
	Handler(route fiber.Router)
}
